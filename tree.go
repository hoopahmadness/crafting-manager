package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	List     CraftingList
	Root     *CraftingItem
	Branches [][]*Tree
}

func newTree(list CraftingList, itemName string) (t Tree, OK bool) {
	t = Tree{}
	t.List = list
	item, OK := list.getItem(itemName)
	t.Root = &item
	return
}

func (this *Tree) graft(other Tree, layer int) {
	if len(this.Branches) == 0 {
		fmt.Println("This shouldn't be happening")
		this.addChoiceLayer()
	}
	this.Branches[layer] = append(this.Branches[layer], &other)
}

func (this *Tree) addChoiceLayer() {
	this.Branches = append(this.Branches, []*Tree{})
}

func (this *Tree) walk() {
	for layerCount, recipe := range this.Root.Recipes {
		// fmt.Println("printing recipe:")
		// fmt.Println(recipe)
		this.addChoiceLayer()
		for _, ingredient := range recipe.Ingredients {
			// fmt.Println("printing ingredient:")
			// fmt.Println(ingredient)
			ingredientName := ingredient.Name
			branch, _ := newTree(this.List, ingredientName)
			branch.walk()
			this.graft(branch, layerCount)
		}
	}
}

func (this Tree) getElements(neededAmount int, resolve map[string]int) (BranchingReport, int) {
	if resolve == nil {
		resolve = map[string]int{}
	}
	forks := 0
	var out BranchingReport
	name := this.Root.Name
	processedRecipes := 0
	for layerCount, recipe := range this.Root.Recipes {
		val, OK := resolve[strings.ToLower(name)]
		if OK && val != layerCount+1 { //check to see if we have a resolution for this branch, and only allow that layer to run
			continue
		}
		processedRecipes++ //this helps us keep track of forks after accounting for resolutions.
		if processedRecipes > 1 {
			forks++
		}
		multiplier := 1
		outputAmount := recipe.Output
		for outputAmount*multiplier < neededAmount {
			multiplier++
		}
		outputAmount = outputAmount * multiplier
		for ingNumber, ing := range recipe.Ingredients {
			inputAmount := ing.Amount * multiplier
			ingredientName := ing.Name
			formatStr := "Craft %d %s from %d %s;"
			nextLevelReport, nextLevelForks := this.Branches[layerCount][ingNumber].getElements(inputAmount, resolve)
			if len(nextLevelReport.Lines) == 0 { //appending a new line to end of every long string
				nextLevelReport.Lines = append(nextLevelReport.Lines, []string{})
			}
			forks += nextLevelForks
			nextLevelReport.insertString(fmt.Sprintf(formatStr, outputAmount, name, inputAmount, ingredientName))
			out.combineReports(nextLevelReport)

		}
		if layerCount < len(this.Root.Recipes)-1 && !OK { //if we're running a resolved layer then don't add the OR
			out.addOR()
		}
	}
	return out, forks
}

func report(litmus, varString, toPrint interface{}) {
	if litmus == varString {
		fmt.Print("REPORT: ")
		fmt.Println(toPrint)
	}
}

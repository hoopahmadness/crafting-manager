package main

import (
	"fmt"
	"strings"
)

var defaultColorRange = []int{-80, 80, -80, 80, 1}

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

func (this Tree) getElements(neededAmount int, resolve map[string]int, colorRange []int) (BranchingReport, int, int) {
	if len(colorRange) == 0 {
		colorRange = defaultColorRange
	}
	if resolve == nil {
		resolve = map[string]int{}
	}
	forks := 0
	// totalRecursions := colorRange[4]
	var out BranchingReport
	name := this.Root.Name
	processedRecipes := 0
	layerDivides := splitRange(colorRange[0], colorRange[1], len(this.Root.Recipes))
	for layerCount, recipe := range this.Root.Recipes {
		layerRange := []int{layerDivides[layerCount], layerDivides[layerCount+1]}
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
		ingredientDivides := splitRange(layerRange[0], layerRange[1], len(recipe.Ingredients))
		for ingNumber, ing := range recipe.Ingredients {
			ingredientRange := []int{ingredientDivides[ingNumber], ingredientDivides[ingNumber+1]}
			inputAmount := ing.Amount * multiplier
			ingredientName := ing.Name
			formatStr := "Craft #%d#^%d %s^ from #%d#^%d %s^;"
			nextColorRange := append([]int{}, ingredientRange...)
			nextColorRange = append(nextColorRange, colorRange[2:4]...)
			nextColorRange = append(nextColorRange, colorRange[4]+1)
			nextLevelReport, nextLevelForks, _ := this.Branches[layerCount][ingNumber].getElements(inputAmount, resolve, nextColorRange)
			if len(nextLevelReport.Lines) == 0 { //appending a new line to end of every long string
				nextLevelReport.Lines = append(nextLevelReport.Lines, []string{})
			}
			forks += nextLevelForks
			// if numIncursions > totalRecursions {
			// 	totalRecursions = numIncursions
			// }
			nextLevelReport.insertString(fmt.Sprintf(formatStr, colorRange[0], outputAmount, name, ingredientRange[0], inputAmount, ingredientName))
			out.combineReports(nextLevelReport)

		}
		if layerCount < len(this.Root.Recipes)-1 && !OK { //if we're running a resolved layer then don't add the OR
			out.addOR()
		}
	}
	return out, forks, -1 //totalRecursions
}

func report(litmus, varString, toPrint interface{}) {
	if litmus == varString {
		fmt.Print("REPORT: ")
		fmt.Println(toPrint)
	}
}

func splitRange(top, bottom, divisions int) []int {
	out := []int{}
	if divisions == 0 {
		return out
	}
	section := (top - bottom) / divisions
	for ii := 0; ii <= divisions; ii++ {
		out = append(out, top-section*ii)
	}
	return out
}

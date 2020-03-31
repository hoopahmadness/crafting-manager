package main

import (
	"fmt"
	"strings"
)

const (
	ORSIGNAL = "OR\n"
)

type Tree struct {
	List     CraftingList
	Root     *CraftingItem
	Branches [][]*Tree
}

func newTree(list CraftingList, itemName string) Tree {
	t := Tree{}
	t.List = list
	item := list[itemName]
	t.Root = &item
	return t
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
			branch := newTree(this.List, ingredientName)
			branch.walk()
			this.graft(branch, layerCount)
		}
	}
}

func (this Tree) getElements(neededAmount int, resolve map[string]int) []string {
	var out []string
	name := this.Root.Name
	for layerCount, recipe := range this.Root.Recipes {
		val, OK := resolve[name]
		if OK && val != layerCount+1 { //check to see if we have a resolution for this branch, and only allow that layer
			continue
		}
		// report(1, layerCount, "On the second layer")
		multiplier := 1
		outputAmount := recipe.Output
		for outputAmount*multiplier < neededAmount {
			multiplier++
		}
		outputAmount = outputAmount * multiplier
		for ingNumber, ing := range recipe.Ingredients {
			inputAmount := ing.Amount * multiplier
			ingredientName := ing.Name
			formatStr := "Craft %d %s from %d %s; %s"
			nextLevelArr := this.Branches[layerCount][ingNumber].getElements(inputAmount, resolve)
			// report(1, layerCount, nextLevelArr)
			if len(nextLevelArr) == 0 {
				nextLevelArr = append(nextLevelArr, "\n")
			}
			for _, nextLevel := range nextLevelArr {
				var outStr string
				templateStr := fmt.Sprintf(formatStr, outputAmount, name, inputAmount, ingredientName, `%s`)
				if nextLevel == ORSIGNAL {
					outStr = strings.Repeat(" ", len(templateStr)) + ORSIGNAL
				} else {
					outStr = fmt.Sprintf(templateStr, nextLevel)
				}
				out = append(out, outStr)
			}
		}
		if layerCount < len(this.Root.Recipes)-1 && !OK { //if we're running a resolved layer then don't add the OR
			out = append(out, ORSIGNAL)
		}
	}
	return out
}

func report(litmus, varString, toPrint interface{}) {
	if litmus == varString {
		fmt.Print("REPORT: ")
		fmt.Println(toPrint)
	}
}
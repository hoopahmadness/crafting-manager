package main

import (
	"encoding/json"
	"fmt"
)

type CraftingList map[string]CraftingItem

func (this CraftingList) addRecipe(name, description string, output int, ingredientList []Ingredient) {
	//add the recipe
	newCraftingItem := newItem(name, description)
	newCraftingItem.Recipes = []Recipe{Recipe{output, ingredientList}}
	this.insertItem(newCraftingItem)

	//add or update the ingredients
	for _, anIngredient := range ingredientList {
		newIngredient := anIngredient.toItem()
		newIngredient.UsedIn = []string{name}
		this.insertItem(newIngredient)
	}

}

//only to be used when inserting a single new recipe and all its ingredients into the crafting list
func (this CraftingList) insertItem(item CraftingItem) {
	oldItem, OK := this[item.Name]
	if !OK {
		this[item.Name] = item
	} else {
		if item.Description != "" {
			oldItem.Description = item.Description
		}
		oldItem.Recipes = append(oldItem.Recipes, item.Recipes...)
		if len(item.UsedIn) != 0 {
			if !contains(oldItem.UsedIn, item.UsedIn[0]) {
				oldItem.UsedIn = append(oldItem.UsedIn, item.UsedIn...)
			}
		}
		this[item.Name] = oldItem
	}
}

func (this CraftingList) listRecipes(itemName string) []Recipe {
	return this[itemName].Recipes
}

func (this CraftingList) listUses(itemName string, recursion bool) []string {
	NUses := this[itemName].UsedIn
	if recursion && len(NUses) > 0 {
		ARC := []string{}
		for _, NUse := range NUses {
			ARC2 := []string{}
			NPlusOneUses := this.listUses(NUse, recursion)
			for _, NPlusOneUse := range NPlusOneUses {
				use := fmt.Sprintf("%s -> %s", NUse, NPlusOneUse)
				ARC2 = append(ARC2, use)
			}
			if len(ARC2) == 0 {
				ARC = append(ARC, NUse)
			} else {
				ARC = append(ARC, ARC2...)
			}
		}
		return ARC
	}
	return NUses
}

func (this CraftingList) listElements(itemName string, itemNumber int, resolutions map[string]int) []string {
	if resolutions == nil {
		resolutions = map[string]int{}
	}
	tree := newTree(this, itemName)
	tree.walk()
	return tree.getElements(itemNumber, resolutions)
}

func (this CraftingList) updateDescription(itemName, description string) {
	item := this[itemName]
	item.Description = description
	this[itemName] = item
}

func (this CraftingList) toBytes() []byte {
	list := []CraftingItem{}
	for _, item := range this {
		if len(item.Recipes) > 0 {
			list = append(list, item)
		}
	}
	b, _ := json.Marshal(list)
	return b
}

func (this CraftingList) fromBytes(b []byte) (err error) {
	list := []CraftingItem{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		return
	}
	for _, item := range list {
		for _, recipe := range item.Recipes {
			this.addRecipe(item.Name, item.Description, recipe.Output, recipe.Ingredients)
		}
	}
	return
}

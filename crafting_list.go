package main

import (
	"encoding/json"
	"fmt"
	"strings"
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
	key := strings.ToLower(item.Name)
	oldItem, OK := this[key]
	if !OK {
		this[key] = item
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
		this[key] = oldItem
	}
}

func (this CraftingList) listRecipes(itemName string) []Recipe {
	itemName = strings.ToLower(itemName)
	return this[itemName].Recipes
}

func (this CraftingList) getItem(itemName string) (item CraftingItem, OK bool) {
	item, OK = this[strings.ToLower(itemName)]
	return
}

func (this CraftingList) listUses(itemName string, recursion bool) []string {
	itemName = strings.ToLower(itemName)
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

func (this CraftingList) getElementTree(itemName string) (Tree, bool) {
	itemName = strings.ToLower(itemName)
	tree, OK := newTree(this, itemName)
	if OK {
		tree.walk()
	}
	return tree, OK
}

func (this CraftingList) updateDescription(itemName, description string) {
	itemName = strings.ToLower(itemName)
	item := this[itemName]
	item.Description = description
	this[itemName] = item
}

func (this CraftingList) toBytes() []byte {
	list := []CraftingItem{}
	for _, item := range this {
		if len(item.Recipes) > 0 || item.Description != "" {
			list = append(list, item)
		}
	}
	b, _ := json.MarshalIndent(list, "", "\t")
	return b
}

//this is gonnaa be a problem later for items with no recipes but with descriptions!!
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

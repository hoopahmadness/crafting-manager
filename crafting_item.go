package main

import (
	"fmt"
)

type CraftingItem struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Recipes     []Recipe `json:"recipes"`
	UsedIn      []string `json:"uses"`
}

func newItem(name, description string) (item CraftingItem) {
	item.Name = name
	item.Description = description
	return
}

type Recipe struct {
	Output      int          `json:"Output"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func (this Ingredient) toItem() CraftingItem {
	return newItem(this.Name, "")
}

func (this CraftingItem) String() string {
	descTemplate := "%s\n"
	thisTemplate := "%s\n%s"
	var descString string
	if this.Description != "" {
		descString = fmt.Sprintf(descTemplate, this.Description)
	} //else leave it blank
	outStr := fmt.Sprintf(thisTemplate, this.Name, descString)

	for _, recipe := range this.Recipes {
		outStr = outStr + recipe.String(this.Name)
	}
	if len(this.UsedIn) > 0 {
		outStr = outStr + "Can be used in: "
		length := len(this.UsedIn)
		for index, use := range this.UsedIn {
			useStr := use
			if index < length-1 {
				useStr = useStr + ","
			}
			outStr = outStr + useStr + " "
		}
	}
	return outStr
}

func (this Recipe) String(name string) string {
	var outStr string
	length := len(this.Ingredients)
	for index, ing := range this.Ingredients {
		outStr = outStr + fmt.Sprintf("%d %s", ing.Amount, ing.Name)
		if index < length-1 {
			outStr = outStr + " + "
		}
	}
	outStr = outStr + fmt.Sprintf(" --> %d %s\n", this.Output, name)
	return outStr
}

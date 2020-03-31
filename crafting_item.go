package main

type CraftingItem struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Recipes     []Recipe `json:"recipes"`
	UsedIn      []string
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

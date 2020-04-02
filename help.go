package main

func generateHelp() (ourMap map[string]string) {
	ourMap = map[string]string{}
	ourMap[USES] = `Prints a list of all the items that can be crafted with this item.
Ex.		
	'> `+USES+` Wax'
			Candle
			Car Wax`

	ourMap[USEALL] = `Prints a branching list of all the items that can be crafted by this item
and the items that be crafted from those, etc. (The inverse of `+ELEMENTS+`.) Use it to see the max potential of
any given item.
Ex.		
	'> `+USEALL+` Tree Branches'
	Tree Branches -> Flimsy Net -> Net -> Gold Net
	Tree Branches -> Flimsy Fishing Rod -> Fishing Rod -> Gold Fishing Rod
	Tree Branches -> Flimsy Axe -> Stone Axe
	Tree Branches -> Flimsy Axe -> Gold Axe
	Tree Branches -> Campfire`

	ourMap[SHOW] = `Shows the name and description for a given item, as well as all known recipes to create it.
Ex.		
	'> `+SHOW+` Net'
	(example to be determined)`

	ourMap[ELEMENTS] = `Starts the Element Wizard to print a branching list of all the ingredients needed to craft a given item.
This tool allows you to see the name and number of every ingredient you need. To get this info, read the end of each line.

In some instances, the item you're polling against or an item further down the tree will have more than one known 
crafting recipe, representing a fork. These are represented in the branching list as a lone 'OR' between two groups.
The OR will more or less line up with the item that is being split. Above the OR will be one or more lines 
describing one recipe for that item, and below will be one or more lines for the second recipe. It's not pretty but it's consistent.

The wizard will ask if you want to choose between forking recipes in order to simplify the output. Simply follow the prompts or exit
to the main menu
Ex.		
	'> `+ELEMENTS+`'
	(follow on-screen prompts to continue)
	...
	Craft 2 Gold Net from 6 Net; Craft 10 Net from 20 Brown Nugget; 
     	                           OR
	Craft 2 Gold Net from 6 Net; Craft 6 Net from 3 Flimsy Net; Craft 3 Flimsy Net from 15 Tree Branches; 
 	Craft 2 Gold Net from 6 Net; Craft 6 Net from 3 Iron Nugget; 
 	Craft 2 Gold Net from 8 Gold Nugget; 
 	OR
 	Craft 3 Gold Net from 8 Golden Spatula;`

	ourMap[INSERT] = `Begins the Recipe Wizard to add new recipes to your crafting list. When this program exits
the new recipes will be saved for later. Carefully read the prompts and respond to the questions.
Ex.		
	'> `+INSERT+`'
	(follow on-screen prompts to continue)`

	ourMap[LISTALL] = `Lists the names of all crafting items that are listed in recipes or as ingredients on your crafting list.
Ex.		
	'> `+LISTALL+`'
	Flimsy Net
	Iron Nugget
	Wood
	Gold Fishing Rod
	Campfire
	Volt Fruit
	Fresh Milk
	Gold Net
	Fishing Rod`

	return
}
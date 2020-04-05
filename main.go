package main

/*TO DO
Add colors to responses (HELP statements, general statements, output from queries, direction, question prompts, emphasis)
add slight delay to responses
Prettier output for some data structures (recipes,)
Add a "SAVE" function to save before closing
Add a "READ" function to add more recipes during operations
Alphabet sort the help menus and listall
split ingredients with |
single-recipe circular logic (catalysts)
Some sort of sorting function so that longest branches are closer to the bottom for a given set of ingredients or for a given set of recipes
*/
import (
	"fmt"
	// "gopkg.in/gookit/color.v1"
	"io/ioutil"
	"os"
	"strings"
)

const (
	//commands
	USES     = "USES"
	USEALL   = "USEALL"
	SHOW     = "SHOW"
	INSERT   = "INSERT"
	ELEMENTS = "ELEMENTS"
	DESCRIBE = "DESCRIBE"
	LISTALL  = "LISTALL"
	HELP     = "HELP"
	EXIT     = "EXIT"
	//replies
	BACK = "BACK"
	YES  = "YES"
	NO   = "NO"
	//errors
	ITEMNOTFOUND = "Item with that name not found"
)

var say Saying

func main() {
	say = renderColors() //get our output object ready
	if len(os.Args) == 1 {
		fmt.Println("Please include an input filename when calling this program.\nFor example, './crafting-manager MyGameCraftingList.json'")
		return
	}
	inputFilePath := os.Args[1]
	jsonBytes, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			//let's just create the file and move on
			print(say.WARNING("That file doesn't exist, so I'll initialize an empty Crafting List and save to that file later!"))
			print(say.HELP("Try using the " + INSERT + " command to add recipes to the new file, or " + HELP + " for more options."))
		} else {
			print(say.DEFAULT(err.Error()))
			return
		}
	}

	list := CraftingList{}
	_ = list.fromBytes(jsonBytes)

	mainLoop(list)
	err = ioutil.WriteFile(inputFilePath, list.toBytes(), 0644) //not sure about these permissions, but why not.
}

func mainLoop(list CraftingList) {
	print(say.DEFAULT("Main Menu!\n"))
	keepRunning := true
	for keepRunning {
		command, arguments := getCommandAndArgs()
		switch command {
		case USES:
			useList := list.listUses(arguments, false)
			print(say.RESULT(strings.Join(useList, "\n")))
		case USEALL:
			useList := list.listUses(arguments, true)
			print(say.RESULT(strings.Join(useList, "\n")))
		case SHOW:
			item, OK := list.getItem(arguments)
			if !OK {
				print(say.DEFAULT(ITEMNOTFOUND))
			} else {
				print(say.RESULT(item))
			}
		case INSERT:
			insertRecipeLoop(list)
		case ELEMENTS:
			getElementsLoop(list)
		case DESCRIBE:
			describeLoop(list, arguments)
		case LISTALL:
			for _, val := range list {
				print(say.RESULT(val.Name))
			}
		case EXIT:
			keepRunning = false
		case HELP, "":
			printHelp(arguments)
		default:
			print(say.HELP("Sorry, didn't recognize that. Try " + HELP + " if you're lost."))
		}
	}
}

func describeLoop(list CraftingList, itemName string) {
	item, OK := list.getItem(itemName)
	if !OK {
		print(say.DEFAULT(ITEMNOTFOUND))
		return
	}
	print(say.DEFAULT(fmt.Sprintf("Current description of %s is: \n '%s'", say.RESULT(itemName), say.RESULT(item.Description))))
	print(say.QUESTION("Do you want to change it?"))
	reply := getYesOrNo()
	if reply == NO {
		return
	}
	print(say.DIRECTION("Please type a new description. Sending a blank reply will not make any changes."))
	newDesc := parseInput()
	if newDesc != "" {
		list.updateDescription(itemName, newDesc)
	}
	print(say.DEFAULT("Done!"))
}

func getElementsLoop(list CraftingList) {
	print(say.QUESTION("What item do you want to break down into its basic elements?"))
	itemName := parseInput()
	print(say.QUESTION("How many of these do you want to make?"))
	outputAmount := getInteger()
	numForks := 1
	output := []string{}
	resolutions := map[string]int{}
	elementTree, OK := list.getElementTree(itemName)
	if !OK {
		print(say.DEFAULT(ITEMNOTFOUND))
		return
	}
	elementTree.walk()
	for numForks != 0 {
		output, numForks = elementTree.getElements(outputAmount, resolutions)
		print(say.RESULT(strings.Join(output, "")))
		if numForks > 0 {
			print(say.QUESTION(`Detecting possible forks. Would you like to resolve these?
YES to make a choice between recipes to show, NO to return to main menu.`))
			reply := getYesOrNo()
			if reply == NO {
				return
			}
			print(say.DIRECTION("Type the name of the item near the 'OR' fork that you wish to resolve."))
			resolveName := parseInput()
			print(say.QUESTION(`Which recipe do you want to keep? Count from top to bottom. For example,
'1' is the group above the OR and '2' is below it. `))
			resolveNumber := getInteger()
			print(say.DEFAULT("Resolving..."))
			resolutions[strings.ToLower(resolveName)] = resolveNumber
		}
	}

}

func insertRecipeLoop(list CraftingList) {
	print(say.DIRECTION("Please follow the prompts to insert a new recipe. At the end of the process you will be given the choice to save or discard."))
	for notDoneRecipe := true; notDoneRecipe; {
		print(say.DIRECTION("Press ENTER to make a new recipe. Send 'BACK' to go back to the main menu"))
		isBack := getBackResponse()
		if isBack {
			return
		}
		print(say.QUESTION("What's the name of the item this recipe crafts?"))
		itemName := parseInput()
		print(say.DIRECTION("If this item is missing a description (or if it's new) you can add one now. Leaving it blank will not make any changes."))
		description := parseInput()
		print(say.QUESTION("How many copies of the item will this recipe yield?"))
		outputAmount := getInteger()

		print(say.DIRECTION("OK, first ingredient:"))
		ingredientList := []Ingredient{}
		for notDoneIngredientList := true; notDoneIngredientList; {
			print(say.QUESTION("What's the name of this ingredient?"))
			ing := parseInput()
			print(say.QUESTION("How many of these do you need?"))
			inputAmount := getInteger()
			print(say.QUESTION("Do you want to save this ingredient to the list?"))
			print(say.RESULT(fmt.Sprintf("Name: %s \nAmount: %d ", ing, inputAmount)))
			answer := getYesOrNo()
			if answer == YES {
				ingredientList = append(ingredientList, Ingredient{ing, inputAmount})
			}
			print(say.QUESTION("Do you need to add more ingredients?"))
			answer = getYesOrNo()
			if answer == NO {
				notDoneIngredientList = false
			} else { //if YES then we start this loop over
				print(say.DEFAULT("OK, let's add another."))
			}
		}
		//after adding ingredients
		//TO DO proper report of recipe
		print(say.DEFAULT("Here's what we got:"))
		print(say.RESULT(fmt.Sprintf("Make %d copies of %s with %d ingredients", outputAmount, itemName, len(ingredientList))))
		print(say.QUESTION("Do you want to save this? NO will discard it."))
		answer := getYesOrNo()
		if answer == YES {
			list.addRecipe(itemName, description, outputAmount, ingredientList)
		} else {
			print(say.DEFAULT("OK, nevermind."))
		}
	}

}

func print(words interface{}) {
	fmt.Println(words)
}

func printHelp(arguments string) {
	helpObject := generateHelp()
	arguments = strings.ToUpper(arguments)
	if helpInfo, OK := helpObject[arguments]; OK {
		print(say.HELP(helpInfo))
	} else {
		print(say.HELP("List of main menu commands:"))
		for key, _ := range helpObject {
			print(say.HELP(key))
		}
		print(say.HELP("Type 'Help <Command>' for more information and usage examples"))
	}
}

const testInput = `[
  {
    "name":"Creamy Heart Soup",
    "description":"Legend of Zelda BotW",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Big Radish|Small Radish|Blue Radish",
            "amount":1
          },
          {
            "name":"Hydromelon",
            "amount":1
          },
          {
            "name":"Volt Fruit",
            "amount":1
          },
          {
            "name":"Fresh Milk",
            "amount":1
          }
        ]
      }
    ]
  },
  {
    "name":"Flimsy Net",
    "description":"10 Bugs 100 Bells",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Tree Branches",
            "amount":5
          }
        ]
      }
    ]
  },
  {
    "name":"Gold Net",
    "description":"",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Net",
            "amount":3
          },
          {
            "name":"Gold Nugget",
            "amount":4
          }
        ]
      },
      {
        "Output":3,
        "ingredients":[
          {
            "name":"Golden Spatula",
            "amount":8
          }
        ]
      }
    ]
  },
  {
    "name":"Flimsy Fishing Rod",
    "description":"10 Fish  100 Bells",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Tree Branches",
            "amount":5
          }
        ]
      }
    ]
  },
  {
    "name":"Fishing Rod",
    "description":"30 Fish  600 Bells",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Flimsy Fishing Rod",
            "amount":1
          },
          {
            "name":"Iron Nugget",
            "amount":1
          }
        ]
      }
    ]
  },
  {
    "name":"Gold Fishing Rod",
    "description":"",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Fishing Rod",
            "amount":1
          },
          {
            "name":"Gold Nugget",
            "amount":1
          }
        ]
      }
    ]
  },
  {
    "name":"Flimsy Axe",
    "description":"40 hits  200 Bells",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Tree Branches",
            "amount":5
          },
          {
            "name":"Stone",
            "amount":1
          }
        ]
      }
    ]
  },
  {
    "name":"Stone Axe",
    "description":"560 Bells",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Flimsy Axe",
            "amount":1
          },
          {
            "name":"Wood",
            "amount":3
          }
        ]
      }
    ]
  },
  {
    "name":"Campfire",
    "description":"",
    "recipes":[
      {
        "Output":1,
        "ingredients":[
          {
            "name":"Tree Branches",
            "amount":3
          }
        ]
      }
    ]
  },
  {
    "name":"Net",
    "description":"30 Bugs 600 Bells",
    "recipes":[
      {
        "Output":5,
        "ingredients":[
          {
            "name":"Brown Nugget",
            "amount":10
          }
        ]
      },
      {
        "Output":2,
        "ingredients":[
          {
            "name":"Flimsy Net",
            "amount":1
          },
          {
            "name":"Iron Nugget",
            "amount":1
          }
        ]
      }
    ]
  }
]`

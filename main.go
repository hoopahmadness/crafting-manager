package main

/*TO DO
Add colors to responses (HELP statements, general statements, output from queries, direction, question prompts, emphasis)
add slight delay to responses
Prettier output for some data structures (recipes,)
Add a "SAVE" function to save before closing
*/
import (
	"fmt"
  "strings"
  "os"
  "io/ioutil"
)

const (
  //commands
	USES      = "USES"
	USEALL    = "USEALL"
	SHOW      = "SHOW"
	INSERT    = "INSERT"
	ELEMENTS  = "ELEMENTS"
	DESCRIBE  = "DESCRIBE"
  LISTALL   = "LISTALL"
  HELP      = "HELP"
	EXIT      = "EXIT"
  //replies
  BACK      = "BACK"
  YES       = "YES"
  NO        = "NO"
  //errors
  ITEMNOTFOUND = "Item with that name not found"
)

func main() {
  if len(os.Args) ==1 {
    fmt.Println("Please include an input filename when calling this program.\nFor example, './crafting-manager MyGameCraftingList.json'")
    return 
  }
  inputFilePath := os.Args[1]
  jsonBytes, err := ioutil.ReadFile(inputFilePath)
  if err != nil {
     if strings.Contains(err.Error(), "no such file or directory") {
      //let's just create the file and move on
      say("That file doesn't exist, so I'll initialize an empty Crafting List and save to that file later!")
      say("Use the "+INSERT+" command to add recipes to the new file!")
     } else {
      say(err.Error())
      return
     }
  }

	list := CraftingList{}
	_ = list.fromBytes(jsonBytes)

  mainLoop(list)
  err = ioutil.WriteFile(inputFilePath, list.toBytes(), 0644) //not sure about these permissions, but why not.
}

func mainLoop(list CraftingList) {
  say("Main Menu!\n")
	keepRunning := true
	for keepRunning {
		command, arguments := getCommandAndArgs()
		switch command {
		case USES:
			useList := list.listUses(arguments, false)
      say(strings.Join(useList, "\n"))
		case USEALL:
			useList := list.listUses(arguments, true)
      say(strings.Join(useList, "\n"))
  	case SHOW:
      item, OK := list.getItem(arguments)
      if !OK {
        say(ITEMNOTFOUND)
      } else {
        say(item)
      }
		case INSERT:
      insertRecipeLoop(list)
		case ELEMENTS:
			getElementsLoop(list)
		case DESCRIBE:
			describeLoop(list, arguments)
		case LISTALL:
			for _, val := range list {
				say(val.Name)
			}
		case EXIT:
			keepRunning = false
    case HELP, "":
      printHelp(arguments)
		default:
			say("Sorry, didn't recognize that.")
		}
	}
}

func describeLoop(list CraftingList, itemName string) {
  item, OK := list.getItem(itemName)
  if !OK {
    say(ITEMNOTFOUND)
    return
  }
  say(fmt.Sprintf("Current description of %s is: \n '%s'" , itemName, item.Description))
  say("Do you want to change it?")
  reply := getYesOrNo()
  if reply==NO {return}
  say("Please type a new description. Sending a blank reply will not make any changes.")
  newDesc := parseInput()
  if newDesc != "" {
    list.updateDescription(itemName, newDesc)    
  }
  say("Done!")
}

func getElementsLoop(list CraftingList) {
  say("What item do you want to break down into its basic elements?")
  itemName := parseInput()
  say("How many of these do you want to make?")
  outputAmount := getInteger()
  numForks := 1
  output := []string{}
  resolutions := map[string]int{}
  elementTree := newTree(list, itemName)
  elementTree.walk()
  for numForks != 0 {
    output, numForks = elementTree.getElements(outputAmount, resolutions)
    say(strings.Join(output, "\n"))
    if numForks >0 {
      say(`Detecting possible forks. Would you like to resolve these?
YES to make a choice between recipes to show, NO to return to main menu.`)
      reply := getYesOrNo()
      if reply == NO {
        return
      }
      say("Type the name of the item near the 'OR' fork that you wish to resolve.")
      resolveName:=parseInput()
      say(`Which recipe do you want to keep? Count from top to bottom. For example,
'1' is the group above the OR and '2' is below it. `)
      resolveNumber := getInteger()
      say("Resolving...")
      resolutions[strings.ToLower(resolveName)] = resolveNumber
    }
  }

}

func insertRecipeLoop(list  CraftingList) {
  say("Please follow the prompts to insert a new recipe. At the end of the process you will be given the choice to save or discard.")
  for notDoneRecipe := true; notDoneRecipe; {
    say("Press ENTER to make a new recipe. Send 'BACK' to go back to the main menu")
    isBack := getBackResponse()
    if isBack {return}
    say("What's the name of the item this recipe crafts?")
    itemName := parseInput()
    say("If this item is missing a description (or if it's new) you can add one below. Leave it blank will not make any changes")
    description := parseInput()
    say("How many copies of the item will this recipe yield?")
    outputAmount := getInteger()

    say("OK, first ingredient:")
    ingredientList := []Ingredient{}
    for notDoneIngredientList := true; notDoneIngredientList; {
      say("What's the name of this ingredient?")
      ing := parseInput()
      say("How many of these do you need?")
      inputAmount := getInteger()
      say("Do you want to save this ingredient to the list?")
      say(fmt.Sprintf("Name: %s \nAmount: %d ", ing, inputAmount))
      answer := getYesOrNo()
      if answer == YES {
        ingredientList = append(ingredientList, Ingredient{ing, inputAmount})
      }
      say("Do you need to add more ingredients?")
      answer = getYesOrNo()
      if answer == NO {
        notDoneIngredientList = false
      } else {  //if YES then we start this loop over
        say("OK, let's add another.")
      }
    }
    //after adding ingredients
    //TO DO proper report of recipe
    say("Here's what we got:")
    say(fmt.Sprintf("Make %d copies of %s with %d ingredients", outputAmount, itemName, len(ingredientList)))
    say("Do you want to save this? NO will discard it.")
    answer := getYesOrNo()
    if answer == YES {
      list.addRecipe(itemName, description, outputAmount, ingredientList)
    } else {
      say("OK, nevermind.")
    }    
  }


}

func say(words interface{}) {
  fmt.Println(words)
}

func printHelp(arguments string) {
  helpObject := generateHelp()
  arguments = strings.ToUpper(arguments)
  if helpInfo, OK := helpObject[arguments]; OK {
    say(helpInfo)
  } else {
    say("List of main menu commands:")
    for key, _ := range helpObject {
      say(key)
    }
    say("Type 'Help <Command>' for more information and usage examples")
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

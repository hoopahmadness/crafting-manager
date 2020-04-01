package main

import (
	"fmt"
  "strings"
)

const (
	USES     = "USES"
	USEALL   = "USEALL"
	SHOW     = "SHOW"
	INSERT   = "INSERT"
	ELEMENTS = "ELEMENTS"
	DESCRIBE = "DESCRIBE"
	LISTALL  = "LISTALL"
	EXIT     = "EXIT"
)

func main() {
	//generate list from file
	//MAIN LOOP
	//	USES  <item> -> print listUses(), back to main loop
	//	USESAll  <item> -> print listUses() recursively, back to main loop
	//	SHOW  <item> -> print report with name, description, recipes, etc return to main loop
	//	INSERT  -> go to insert loop, make x number of recipes, return to main loop
	//	ELEMENTS  -> go to insert loop, make x number of recipes, return to main loop
	//	DESCRIBE  <item> -> print current description, then allows user input to replace it, returns to main loop
	//	LISTALL  -> Print all item names in list with descriptions
	//	EXIT -> write to file, stop program
	//

	list := CraftingList{}

	_ = list.fromBytes([]byte(testInput))
  mainLoop(list)
}

func mainLoop(list CraftingList) {
	keepRunning := true
	for keepRunning {
		command, arguments := parseInput()
		switch command {
		case USES:
			useList := list.listUses(arguments, false)
      fmt.Println(strings.Join(useList, "\n"))
		case USEALL:
			useList := list.listUses(arguments, true)
      fmt.Println(strings.Join(useList, "\n"))
  		case SHOW:
			fmt.Println(list[arguments])
		case INSERT:
			//insert wizard loop
		case ELEMENTS:
			//elements loop
		case DESCRIBE:
			//description loop
		case LISTALL:
			for k, _ := range list {
				fmt.Println(k)
			}
		case EXIT:
			keepRunning = false
		default:
			fmt.Println("Sorry, didn't recognize that.")
		}
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

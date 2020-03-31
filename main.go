package main

import (
	"fmt"
	// "os"
	// "github.com/urfave/cli/v2"
)

func main() {
	list := CraftingList{}

	_ = list.fromBytes([]byte(testInput))
	// fmt.Println(list.listUses("Tree Branches", true)[0])
	fmt.Println(list.listElements("Gold Net", 2, map[string]int{"Net": 2}))
	// &cli.App{}).Run(os.Args)

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

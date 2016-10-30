package main

import (
	"fmt"

	"github.com/geeks-dev/dilemma"
)

type Sweets struct {
	Name  string
	Price string
}

func main() {
	fmt.Println()

	{
		sweets := []Sweets{
			Sweets{
				Name:  "waffles",
				Price: "$2",
			},
			Sweets{
				Name:  "ice cream",
				Price: "$1",
			},
			Sweets{
				Name:  "candy",
				Price: "$1",
			},
			Sweets{
				Name:  "biscuits",
				Price: "$1",
			},
		}
		s := dilemma.Config{
			Title:   "Hello there!\n\rSelect a treat using the arrow keys:",
			Help:    "Use arrow up and down, then enter to select.\n\rChoose wisely.",
			Options: sweets,
			Key:     "Name",
		}
		selected, exitKey, err := dilemma.Prompt(s)
		if err != nil || exitKey == dilemma.CtrlC {
			fmt.Print("Exiting...\n")
			return
		}

		fmt.Printf("Enjoy your %s!\n", sweets[selected].Price)
	}

	fmt.Println()

	{
		type Animal struct {
			Name  string
			Color string
		}
		animals := []Animal{
			Animal{
				Name:  "dog",
				Color: "black",
			},
			Animal{
				Name:  "pony",
				Color: "brown",
			},
			Animal{
				Name:  "cat",
				Color: "yellow",
			},
			Animal{
				Name:  "rabbit",
				Color: "white",
			},
			Animal{
				Name:  "gopher",
				Color: "light blue",
			},
			Animal{
				Name:  "elephant",
				Color: "gray",
			},
		}

		s := dilemma.Config{
			Title:   "Do what color do you see?",
			Help:    "Use arrow up and down, then enter to select.",
			Options: animals,
			Key:     "Color",
		}
		selected, exitKey, err := dilemma.Prompt(s)
		if err != nil || exitKey == dilemma.CtrlC {
			fmt.Print("Exiting...\n")
			return
		}

		fmt.Printf("It %s!\n", animals[selected].Name)
	}

	fmt.Println()
}

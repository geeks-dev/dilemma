package main

import (
	"fmt"

	"github.com/geeks-dev/dilemma"
)

func main() {
	fmt.Println()

	{
		s := dilemma.Config{
			Title: "Hello there!\n\rSelect a treat using the arrow keys:",
			Help:  "Use arrow up and down, then enter to select.\n\rChoose wisely.",
			Options: []map[string]string{
				{
					"name":  "waffles",
					"price": "$2",
				},
				{
					"name":  "ice cream",
					"price": "$1",
				},
				{
					"name":  "candy",
					"price": "$1",
				},
				{
					"name":  "biscuits",
					"price": "$1",
				},
			},
			Key: "name",
		}
		selected, exitKey, err := dilemma.Prompt(s)
		if err != nil || exitKey == dilemma.CtrlC {
			fmt.Print("Exiting...\n")
			return
		}

		fmt.Printf("Enjoy your %s!\n", selected)
	}

	fmt.Println()

	{
		s := dilemma.Config{
			Title: "Do what color do you see?",
			Help:  "Use arrow up and down, then enter to select.",
			Options: []map[string]string{
				{
					"name":  "dog",
					"color": "black",
				},
				{
					"name":  "pony",
					"color": "brown",
				},
				{
					"name":  "cat",
					"color": "yellow",
				},
				{
					"name":  "rabbit",
					"color": "white",
				},
				{
					"name":  "gopher",
					"color": "light blue",
				},
				{
					"name":  "elephant",
					"color": "gray",
				},
			},
			Key: "color",
		}
		selected, exitKey, err := dilemma.Prompt(s)
		if err != nil || exitKey == dilemma.CtrlC {
			fmt.Print("Exiting...\n")
			return
		}

		fmt.Printf("It %s!\n", selected["name"])
	}

	fmt.Println()
}

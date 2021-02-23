package main

import "fmt"

func main() {
	/*
		// Assign values to a map
			var colors map[string]string

			colors := make(map[string]string)
			colors["white"] = "#FFFFFF"
			colors[10]= "#FFFFFF"

		// Delere values from a map
			delete(colors, 10)
	*/
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

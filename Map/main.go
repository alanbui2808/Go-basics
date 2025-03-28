package main

import "fmt"

func main() {
	// Note: Map is reference type
	// Struct is value type

	// map of key string with value string (hex code)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#45fb34",
	}

	// create an empty map
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color:", color, "with hexcode:", hex)
	}
}

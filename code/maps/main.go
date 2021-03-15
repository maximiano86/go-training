package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "sun",
		"blue": "sky",
	}

	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "snow"
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, word := range c {
		fmt.Println(color, word)
	}
}

package main

import "fmt"

func main() {
	colors2 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)

	printMap(colors2)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("key:", color, "value:", hex)
	}
}

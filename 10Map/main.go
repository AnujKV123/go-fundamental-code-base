package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	colors["yellow"] = "#ffff00"

	fmt.Println("List of all the colors: ", colors)
	fmt.Println("white color: ", colors["white"])

	delete(colors, "black")
	fmt.Println("List of all the colors: ", colors)

	// loops through the map
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

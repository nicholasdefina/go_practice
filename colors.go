package main

import "fmt"

// practice go maps

func createColorsMap() map[string]string {
	colors := map[string]string{ // [key]value (must be uniform - ie. all keys are one type, all vals are one type)
		"red":    "#ff0000",
		"blue":   "#0000ff",
		"yellow": "#ffff00",
	}
	return colors

}

func printColors(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex for", key, "is", val)
	}

}

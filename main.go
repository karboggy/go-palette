package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cascax/colorthief-go"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <image-path> <color-count>", os.Args[0])
	}
	imagePath := os.Args[1]
	colorCount, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// generate palette
	palette, err := colorthief.GetPaletteFromFile(imagePath, colorCount)
	if err != nil {
		log.Fatal(err)
	}

	// print palette as hex colors
	for _, c := range palette {
		r, g, b, _ := c.RGBA()
		fmt.Printf("#%02x%02x%02x ", uint8(r>>8), uint8(g>>8), uint8(b>>8))
	}
	fmt.Printf("\n")
}

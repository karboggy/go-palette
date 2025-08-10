package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cascax/colorthief-go"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <image-path>", os.Args[0])
	}
	imagePath := os.Args[1]

	// generate palette
	palette, err := colorthief.GetPaletteFromFile(imagePath, 16)
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

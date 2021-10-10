package main

import (
	"fmt"
	"minesweeper/field"
	"minesweeper/input"
)

func main() {
	f := field.NewField(8)
	f.Display()

	for {
		f.Select(input.GetInput())
		f.Display()
		fmt.Printf("\n")
	}
}

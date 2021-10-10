package main

import (
	"fmt"
	"minesweeper/field"
	"minesweeper/input"
)

func main() {
	f := field.NewField(8)
	f.Display()

	fmt.Printf("%v\n", input.GetInput())
}

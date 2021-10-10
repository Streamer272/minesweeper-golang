package main

import (
	"fmt"
	"minesweeper/field"
	"minesweeper/input"
)

func main() {
	f := field.NewField(8)

	for {
		f.Display()

		key := input.GetInput()
		if key == input.UP || key == input.LEFT || key == input.DOWN || key == input.RIGHT {
			f.Select(key)
		} else if key == input.UNCOVER {

		} else if key == input.FLAG {

		} else {
			fmt.Printf("Wrong keypress")
		}

		fmt.Printf("\n")
	}
}

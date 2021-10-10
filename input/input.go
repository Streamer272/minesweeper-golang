package input

import (
	tty "github.com/mattn/go-tty"
)

func GetInput() rune {
	reader, _ := tty.Open()
	defer reader.Close()

	r, _ := reader.ReadRune()
	return r
}

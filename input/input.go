package input

import (
	"github.com/mattn/go-tty"
)

const (
	UP = iota
	LEFT
	DOWN
	RIGHT

	UNCOVER
	FLAG
)

const (
	w = 119
	a = 97
	s = 115
	d = 100

	space = 32
	f     = 102
)

func GetInput() int {
	reader, _ := tty.Open()
	defer reader.Close()

	r, _ := reader.ReadRune()
	return GetAction(int(r))
}

func GetAction(r int) int {
	switch r {
	case w:
		return UP
	case a:
		return LEFT
	case s:
		return DOWN
	case d:
		return RIGHT

	case space:
		return UNCOVER
	case f:
		return FLAG

	default:
		return 0
	}
}

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
	W = 119
	A = 97
	S = 115
	D = 100

	SPACE = 32
	F     = 102
)

func GetInput() int {
	reader, _ := tty.Open()
	defer reader.Close()

	r, _ := reader.ReadRune()
	return GetAction(int(r))
}

func GetAction(r int) int {
	switch r {
	case W:
		return UP
	case A:
		return LEFT
	case S:
		return DOWN
	case D:
		return RIGHT

	case SPACE:
		return UNCOVER
	case F:
		return FLAG

	default:
		return -1
	}
}

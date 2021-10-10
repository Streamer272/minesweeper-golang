package input

import (
	"github.com/mattn/go-tty"
)

const (
	NIL = iota
	UP
	LEFT
	DOWN
	RIGHT
)

var (
	w = 119
	a = 97
	s = 115
	d = 100

	h = 104
	j = 106
	k = 107
	l = 108

	k1 = 49
	k2 = 50
	k3 = 51
	k4 = 52

	upArr    = []int{w, k, k1}
	leftArr  = []int{a, h, k4}
	downArr  = []int{s, j, k3}
	rightArr = []int{d, l, k2}
)

func GetInput() int {
	reader, _ := tty.Open()
	defer reader.Close()

	r, _ := reader.ReadRune()
	return GetDirectionByInput(int(r))
}

func GetDirectionByInput(r int) int {
	for k := range upArr {
		if upArr[k] == r {
			return UP
		}
	}

	for k := range leftArr {
		if k == r {
			return LEFT
		}
	}

	for k := range downArr {
		if k == r {
			return DOWN
		}
	}

	for k := range rightArr {
		if k == r {
			return RIGHT
		}
	}

	return NIL
}

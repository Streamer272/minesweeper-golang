package field

import (
	"minesweeper/field/box"
	"minesweeper/input"
)

func (f *Field) Select(direction int) {
	var selected = 0
	for selected = range f.Boxes {
		if f.Boxes[selected].Selected {
			break
		}
	}

	switch direction {
	case input.UP:
		if selected < f.Size {
			return
		}
	case input.LEFT:
		if selected%f.Size == 0 {
			return
		}
	case input.DOWN:
		if selected >= f.Size*(f.Size-1) {
			return
		}
	case input.RIGHT:
		if selected%f.Size == f.Size-1 {
			return
		}
	default:
		return
	}

	switch direction {
	case input.UP:
		f.Boxes[selected-f.Size].Selected = true
	case input.LEFT:
		f.Boxes[selected-1].Selected = true
	case input.DOWN:
		f.Boxes[selected+f.Size].Selected = true
	case input.RIGHT:
		f.Boxes[selected+1].Selected = true
	}

	f.Boxes[selected].Selected = false
}

func (f *Field) Uncover() {
	for selected := range f.Boxes {
		if f.Boxes[selected].Selected {
			f.Boxes[selected].State = box.VISIBLE
			return
		}
	}
}

func (f *Field) Flag() bool {
	for selected := range f.Boxes {
		if f.Boxes[selected].Selected {
			f.Boxes[selected].State = box.FLAGGED
			return f.Boxes[selected].Value != box.BOMB
		}
	}

	return true
}

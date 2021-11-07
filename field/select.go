package field

import (
	"minesweeper/field/box"
	"minesweeper/input"
)

func (f *Field) Select(direction int) {
	selected := f.selected()

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

func (f *Field) UncoverSelected() bool {
	if f.IsEmpty() {
		f.Init()
	}

	selected := f.selected()
	f.Boxes[selected].State = box.VISIBLE

	var uncovered []int
	f.UncoverSurrounding(selected, &uncovered)

	return f.Boxes[selected].Value != box.BOMB
}

func (f *Field) UncoverSurrounding(selected int, uncovered *[]int) {
	for _, uncoveredIndex := range *uncovered {
		if selected == uncoveredIndex {
			return
		}
	}

	if f.Boxes[selected].Value == box.BOMB || f.getSurroundingBombCount(selected) != 0 {
		return
	}

	if f.Boxes[selected].Value != box.EMPTY {
		*uncovered = append(*uncovered, selected)
		return
	}

	f.Boxes[selected].State = box.VISIBLE

	for _, dir := range f.Dir {
		if f.isInBounds(selected + dir) {
			f.UncoverSurrounding(selected+dir, uncovered)
		}
	}
}

func (f *Field) Flag() {
	selected := f.selected()
	f.Boxes[selected].State = box.FLAGGED
}

func (f *Field) selected() int {
	for _, b := range f.Boxes {
		if b.Selected {
			return b.Index
		}
	}

	return -1
}

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

func (f *Field) UncoverSelected() bool {
	if f.IsEmpty() {
		f.Init()
	}

	var selected int
	for selected = range f.Boxes {
		if f.Boxes[selected].Selected {
			f.Boxes[selected].State = box.VISIBLE

			break
		}
	}

	//var uncovered *[]int = &[]int{}
	//f.UncoverSurrounding(selected, uncovered)

	return f.Boxes[selected].Value != box.BOMB

	// TODO uncover while no bombs around
}

func (f *Field) UncoverSurrounding(index int, uncovered *[]int) {
	// FIXME
	//fmt.Printf("Called uncover surrounding with args %v %v\n", index, *uncovered)

	//time.Sleep(time.Second)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			currentIndex := index + f.Size*i + j
			//fmt.Printf("on index %v\n", currentIndex)

			if currentIndex == index || currentIndex < 0 || currentIndex >= f.Size*f.Size {
				continue
			}

			for _, uncoveredIndex := range *uncovered {
				//fmt.Printf("Checking skipping (%v == %v)...\n", currentIndex, uncoveredIndex)
				if currentIndex == uncoveredIndex {
					//time.Sleep(time.Millisecond * 500)
					continue
				}
			}

			if f.Boxes[currentIndex].Value != box.BOMB {
				//fmt.Printf("%v was not a bomb\n", currentIndex)

				f.Boxes[currentIndex].State = box.VISIBLE
				//*uncovered = append(*uncovered, index)

				if f.getSurroundingBombCount(currentIndex) == 0 {
					f.UncoverSurrounding(currentIndex, uncovered)
				}

				//f.Display()

				//time.Sleep(time.Second)
			}
		}
	}
}

func (f *Field) Flag() {
	for selected := range f.Boxes {
		if f.Boxes[selected].Selected {
			f.Boxes[selected].State = box.FLAGGED
			return
		}
	}
}

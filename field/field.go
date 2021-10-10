package field

import (
	"fmt"
	"math"
	"minesweeper/field/box"
	"minesweeper/input"
)

type Field struct {
	Size  int
	Boxes []box.Box
}

func NewField(size int) Field {
	if size < 5 {
		panic("Size cannot be smaller than 5")
	}

	f := Field{
		Size: size,
	}

	f.Boxes = []box.Box{}

	for i := 0; i < size*size; i++ {
		f.Boxes = append(f.Boxes,
			box.NewBox(i, i == (f.Size*int(math.Floor(float64(f.Size/2))))-1-int(math.Floor(float64(f.Size/2)))+f.Size),
		)
	}

	return f
}

func (f *Field) Display() {
	for i := 0; i < f.Size*f.Size; i++ {
		if i%f.Size == 0 {
			if f.Boxes[i].Selected {
				fmt.Printf("[ %v", f.Boxes[i].AsSymbol())
			} else {
				fmt.Printf("| %v", f.Boxes[i].AsSymbol())
			}
		} else if i%f.Size == f.Size-1 {
			if f.Size%2 == 0 {
				if f.Boxes[i].Selected {
					fmt.Printf(" [ ")
				} else if f.Boxes[i-1].Selected {
					fmt.Printf(" ] ")
				} else {
					fmt.Printf(" | ")
				}
			}

			if f.Boxes[i].Selected {
				fmt.Printf("%v ]\n", f.Boxes[i].AsSymbol())
			} else {
				fmt.Printf("%v |\n", f.Boxes[i].AsSymbol())
			}
		} else if (i%f.Size)%2 == 1 {
			if f.Boxes[i-1].Selected {
				fmt.Printf(" ] %v | ", f.Boxes[i].AsSymbol())
			} else if f.Boxes[i].Selected {
				fmt.Printf(" [ %v ] ", f.Boxes[i].AsSymbol())
			} else if f.Boxes[i+1].Selected {
				fmt.Printf(" | %v [ ", f.Boxes[i].AsSymbol())
			} else {
				fmt.Printf(" | %v | ", f.Boxes[i].AsSymbol())
			}
		} else {
			fmt.Printf("%v", f.Boxes[i].AsSymbol())
		}
	}
}

func (f *Field) Select(direction int) {
	var currentUncovered = 0
	for currentUncovered = range f.Boxes {
		if f.Boxes[currentUncovered].Selected {
			f.Boxes[currentUncovered].Selected = false
			break
		}
	}

	switch direction {
	case input.UP:
		if currentUncovered < f.Size {
			return
		}
	case input.LEFT:
		if currentUncovered%f.Size == 0 {
			return
		}
	case input.DOWN:
		if currentUncovered > f.Size*f.Size-f.Size {
			return
		}
	case input.RIGHT:
		if currentUncovered%f.Size == f.Size-1 {
			return
		}
	default:
		return
	}

	switch direction {
	case input.UP:
		f.Boxes[currentUncovered-f.Size].Selected = true
	case input.LEFT:
		f.Boxes[currentUncovered-1].Selected = true
	case input.DOWN:
		f.Boxes[currentUncovered+f.Size].Selected = true
	case input.RIGHT:
		f.Boxes[currentUncovered+1].Selected = true
	}
}

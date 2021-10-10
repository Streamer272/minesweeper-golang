package field

import (
	"fmt"
	"math"
	"minesweeper/field/box"
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

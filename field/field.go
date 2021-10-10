package field

import (
	"fmt"
	"github.com/inancgumus/screen"
	"math"
	"minesweeper/field/box"
	"minesweeper/input"
	"os"
)

type Field struct {
	Size  int
	Boxes []box.Box
}

func NewField(size int) Field {
	if size < 5 {
		fmt.Printf("Size cannot be smaller than 5\n")
		os.Exit(1)
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
	width, height := screen.Size()

	if height > f.Size+1 {
		for i := 0; float64(i) < math.Floor(float64((height-f.Size)/2)); i++ {
			fmt.Printf("\n")
		}
	}

	for i := 0; i < f.Size*f.Size; i++ {
		if i%f.Size == 0 {
			if width > f.Size+1 {
				for i := 0; float64(i) < math.Floor(float64((width-(f.Size*4+1))/2)); i++ {
					fmt.Printf(" ")
				}
			}

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

	if height > f.Size+1 {
		for i := 0; float64(i) < math.Floor(float64((height-f.Size)/2)); i++ {
			fmt.Printf("\n")
		}
	}

	if width < f.Size || height < f.Size {
		fmt.Printf("Your window is too small, please make it bigger!")
		fmt.Printf("(current resulution: %vx%v, required resolution: %vx%v)\n", width, height, f.Size, f.Size)
	}
}

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

func (f *Field) Flag() {
	for selected := range f.Boxes {
		if f.Boxes[selected].Selected {
			f.Boxes[selected].State = box.FLAGGED
			return
		}
	}
}

package field

import (
	"fmt"
	"github.com/inancgumus/screen"
	"math"
	"minesweeper/field/box"
	"strconv"
)

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
				fmt.Printf("[ %v", f.asSymbol(i))
			} else {
				fmt.Printf("| %v", f.asSymbol(i))
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
				fmt.Printf("%v ]\n", f.asSymbol(i))
			} else {
				fmt.Printf("%v |\n", f.asSymbol(i))
			}
		} else if (i%f.Size)%2 == 1 {
			if f.Boxes[i-1].Selected {
				fmt.Printf(" ] %v | ", f.asSymbol(i))
			} else if f.Boxes[i].Selected {
				fmt.Printf(" [ %v ] ", f.asSymbol(i))
			} else if f.Boxes[i+1].Selected {
				fmt.Printf(" | %v [ ", f.asSymbol(i))
			} else {
				fmt.Printf(" | %v | ", f.asSymbol(i))
			}
		} else {
			fmt.Printf("%v", f.asSymbol(i))
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

func (f *Field) getSurroundingBombCount(index int) int {
	bombCounter := 0
	if index-f.Size-1 >= 0 && (index-f.Size-1)%f.Size != f.Size - 1 && f.Boxes[index-f.Size-1].Value == box.BOMB {
		bombCounter++
	}
	if index-f.Size >= 0 && f.Boxes[index-f.Size].Value == box.BOMB {
		bombCounter++
	}
	if index-f.Size+1 >= 0 && (index-f.Size+1)%f.Size != 0 && f.Boxes[index-f.Size+1].Value == box.BOMB {
		bombCounter++
	}
	if (index-1)%f.Size != f.Size - 1 && f.Boxes[index-1].Value == box.BOMB {
		bombCounter++
	}
	if (index+1)%f.Size != 0 && f.Boxes[index+1].Value == box.BOMB {
		bombCounter++
	}
	if index+f.Size-1 < f.Size*f.Size && (index+f.Size-1)%f.Size != f.Size - 1 && f.Boxes[index+f.Size-1].Value == box.BOMB {
		bombCounter++
	}
	if index+f.Size < f.Size*f.Size && f.Boxes[index+f.Size].Value == box.BOMB {
		bombCounter++
	}
	if index+f.Size+1 < f.Size*f.Size && (index+f.Size+1)%f.Size != 0 && f.Boxes[index+f.Size+1].Value == box.BOMB {
		bombCounter++
	}

	return bombCounter
}

func (f *Field) asSymbol(index int) string {
	if f.Boxes[index].AsSymbol() != "0" {
		return f.Boxes[index].AsSymbol()
	}

	return strconv.Itoa(f.getSurroundingBombCount(index))
}

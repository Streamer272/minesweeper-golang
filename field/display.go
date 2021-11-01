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
	output := ""

	if height > f.Size+1 {
		for i := 0; float64(i) < math.Floor(float64((height-f.Size)/2)); i++ {
			output += "\n"
		}
	}

	for i := 0; i < f.Size*f.Size; i++ {
		if i%f.Size == 0 {
			if width > f.Size+1 {
				for i := 0; float64(i) < math.Floor(float64((width-(f.Size*4+1))/2)); i++ {
					output += " "
				}
			}

			if f.Boxes[i].Selected {
				output += fmt.Sprintf("[ %v", f.asSymbol(i))
			} else {
				output += fmt.Sprintf("| %v", f.asSymbol(i))
			}
		} else if i%f.Size == f.Size-1 {
			if f.Size%2 == 0 {
				if f.Boxes[i].Selected {
					output += " [ "
				} else if f.Boxes[i-1].Selected {
					output += " ] "
				} else {
					output += " | "
				}
			}

			if f.Boxes[i].Selected {
				output += fmt.Sprintf("%v ]\n", f.asSymbol(i))
			} else {
				output += fmt.Sprintf("%v |\n", f.asSymbol(i))
			}
		} else if (i%f.Size)%2 == 1 {
			if f.Boxes[i-1].Selected {
				output += fmt.Sprintf(" ] %v | ", f.asSymbol(i))
			} else if f.Boxes[i].Selected {
				output += fmt.Sprintf(" [ %v ] ", f.asSymbol(i))
			} else if f.Boxes[i+1].Selected {
				output += fmt.Sprintf(" | %v [ ", f.asSymbol(i))
			} else {
				output += fmt.Sprintf(" | %v | ", f.asSymbol(i))
			}
		} else {
			output += fmt.Sprintf("%v", f.asSymbol(i))
		}
	}

	if height > f.Size+1 {
		for i := 0; float64(i) < math.Floor(float64((height-f.Size)/2)); i++ {
			output += "\n"
		}
	}

	if width < f.Size || height < f.Size {
		output += "Your window is too small, please make it bigger!"
		output += fmt.Sprintf("(current resulution: %vx%v, required resolution: %vx%v)\n", width, height, f.Size, f.Size)
	}

	fmt.Printf("%v", output)
}

func (f *Field) getSurroundingBombCount(index int) int {
	bombCount := 0
	for i := index/f.Size - 1; i <= index/f.Size+1; i++ {
		for j := index%f.Size - 1; j <= index%f.Size+1; j++ {
			if i < 0 || i >= f.Size ||
				j < 0 || j >= f.Size {
				continue
			}
			if f.Boxes[f.Size*i+j].Value == box.BOMB {
				bombCount++
			}

		}
	}

	return bombCount
}

func (f *Field) asSymbol(index int) string {
	if f.Boxes[index].AsSymbol() != "0" {
		return f.Boxes[index].AsSymbol()
	}

	return strconv.Itoa(f.getSurroundingBombCount(index))
}

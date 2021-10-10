package main

import (
	"fmt"
	"minesweeper/field"
	"minesweeper/input"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	f := field.NewField(8)

	clear()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		clear()
		os.Exit(0)
	}()

	for {
		f.Display()

		key := input.GetInput()
		if key == input.UP || key == input.LEFT || key == input.DOWN || key == input.RIGHT {
			f.Select(key)
		} else if key == input.UNCOVER {
			f.Uncover()
		} else if key == input.FLAG {
			f.Flag()
		} else {
			fmt.Printf("Wrong keypress")
		}

		clear()
	}
}

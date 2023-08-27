package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func DisplayMatrix(matrix [][]Cell) {
	displayString := ""

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			switch matrix[i][j].alive {
			case false:
				displayString += "  "
			case true:
				displayString += " ■"
			}
		}
		displayString += "\n"
	}

	clearScreen()
	fmt.Printf(displayString)
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear") //Linux example, its tested
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	default:
		fmt.Println("CLS for ", runtime.GOOS, " not implemented")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

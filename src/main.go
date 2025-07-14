package main

import (
	"GoForge/utils"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Yellow("Usage: goforge [run | build | new <pkg-name> ]\n")
		return
	}
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			color.Yellow("Usage: goforge [run | build | new <pkg-name> ]\n")
			return
		}
		utils.New()
	case "run":
		utils.Run()
	case "build":
		if len(os.Args) == 3 {
			switch os.Args[2] {
			case "run":
				utils.Build()
				utils.Chvenv("../")
				utils.Run()
			default:
				color.Red("Argument '%v' not defined!\n", os.Args[2])
			}
		} else {
			utils.Build()
		}
	case "install":
		utils.Copy()
	default:
		color.Yellow("Usage: goforge [run | build | new <pkg-name> ]\n")
		return
	}
}

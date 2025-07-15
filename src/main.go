package main

import (
	"GoForge/utils"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Yellow("⚠️  Usage: goforge [run | build | init <pkg-name> | install ]\n")
		return
	}
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			color.Yellow("⚠️  Usage: goforge [run | build | init <pkg-name> | install ]\n")
			return
		}
		utils.New()
	case "run":
		utils.Run()
	case "build":
		utils.Buildscr()
	case "install":
		utils.Buildscr()
		utils.Install()
	default:
		color.Yellow("⚠️  Usage: goforge [run | build | new <pkg-name> | install ]\n")
		return
	}
}

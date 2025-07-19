package main

import (
	"GoForge/utils"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Yellow("⚠️  Usage: goforge <command> [arguments]\n")
		return
	}
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			color.Yellow("⚠️  Usage: goforge <command> [arguments]\n")
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
	case "remove":
		utils.Remove()
	case "help":
		fmt.Println(`Goforge - A minimal forge to build and manage your Go-based projects

Usage:
  goforge [command] [arguments]

Available Commands:
  help                 Show this help message
  version              Show the current version of goforge
  run                  Run the current project (main package)
  build                Build the project and output the executable
  new <pkg-name>       Initialize a new goforge project with the given package name
  install              Install project as a program in GOBIN
  remove               Remove the installed program from GOBIN
  clean	               Removes all builds and temporary files

For more information, visit: https://example.com/goforge
`)
	default:
		color.Yellow("⚠️  Usage: goforge <command> [arguments]\n")
		return
	}
}

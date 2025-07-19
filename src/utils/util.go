package utils

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var destFolder = Gobin() // <- change as you like

const CONFIG_FILE = "GoForge.yaml"

var srcfilename = "./src/main.go"

var SrcContent = `package main

import "fmt"

func main(){
	fmt.Println("Project Initialised by GoForge!")
}`

var Pkg string

var cfgcontent = `  optimisation: true
  
  env:
    GOOS: windows
    GOARCH: amd64

  flags:
    - -ldflags
    - "-s -w"`

type Config struct {
	App struct {
		Package string `yaml: "package"`
		Version string `yaml: "version"`
	} `yaml: "app"`

	Build struct {
		Output       string `yaml: "output"`
		Optimisation bool   `yaml: "optimisation"`

		Env   map[string]string `yaml: "env"`
		Flags []string          `yaml: "flags"`
	} `yaml: "build"`
}

func New() {
	CreateCfgFile(CONFIG_FILE, cfgcontent)

	dirname := "./src"

	// Check if directory exists
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		CreateFile(srcfilename, SrcContent)
	} else if err != nil {
		// Some other error occurred
		fmt.Println("Error checking directory:", err)
	} else if !info.IsDir() {
		// Path exists but is not a directory
		fmt.Printf("src is not a directory!\n")
	} else {
		// Directory exists, do nothing

		CreateFile(srcfilename, SrcContent)
	}

	Init()

}

func Build() {
	cfg, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		color.Red("❌ Failed to load config: %v\n", err)
		return
	}

	switch cfg.Build.Optimisation {
	case true:
		for key, val := range cfg.Build.Env {
			os.Setenv(key, val)
		}

		Chvenv("src")
		err = Cmd("go", "mod", "tidy")
		if err != nil {
			color.Red("❌ Error building app, pls check your modules: %v\n", err)
			return
		}
		// making cmd
		args := append(cfg.Build.Flags, "-o", "../"+cfg.Build.Output, "main.go")

		err = Cmd("go", append([]string{"build"}, args...)...)

		color.Blue("🔨 Running build:")
		color.Cyan("go %v\n", args)

		if err != nil {
			color.Red("❌ Build Failed: %v\n", err)
			return
		}
		color.Green("✅ Build Successful!\n")
	case false:
		Chvenv("src")
		err = Cmd("go", "mod", "tidy")
		if err != nil {
			color.Red("❌ Error building app, pls check your modules: %v\n", err)
			return
		}
		err = Cmd("go", "build", "-o", cfg.Build.Output)

		color.Blue("🔨 Running build:")
		color.Cyan("go build -o %v\n", cfg.Build.Output)

		if err != nil {
			color.Red("❌ Build Failed: %v\n", err)
			return
		}
		color.Green("✅ Build Successful!\n")
	}
	Chvenv("../")

}

func Run() {
	cfg, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		color.Red("❌ Failed to load config: %v\n", err)
		return
	}
	if !FileExists(cfg.Build.Output) {
		Build()
		Chvenv("../")
		Run()
		return
	}
	color.Green("--------------------\n")
	color.Blue("Running Program...\n")
	color.Green("--------------------\n")
	fmt.Println()

	err = Cmd(cfg.Build.Output)

	if err != nil {
		color.Red("Error running exe: %v\n", err)
		fmt.Println()
		return
	}
	fmt.Println()
}

func Install() {
	cfg, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		color.Red("❌ Failed to load config: %v\n", err)
		return
	}
	src := cfg.Build.Output

	// Resolve absolute paths
	absSrc, err := filepath.Abs(src)
	Check(err)

	// Verify source exists and is a regular file
	info, err := os.Stat(absSrc)
	Check(err)
	if info.IsDir() {
		fmt.Printf("Error: '%s' is a directory, not a file\n", absSrc)
		os.Exit(1)
	}

	// Ensure destination folder exists
	err = os.MkdirAll(destFolder, 0o755)
	Check(err)

	destPath := filepath.Join(destFolder, filepath.Base(absSrc))

	// Copy (overwrite if exists)
	err = CopyFile(absSrc, destPath)
	Check(err)

	fmt.Printf("File copied to %s\n", destPath)
}

func Remove() {
	cfg, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		color.Red("❌ Failed to load config: %v\n", err)
		return
	}

	binPath := Gobin()
	out := strings.Replace(cfg.Build.Output, "build/", "", 1)
	src := binPath + "/" + out

	err = os.Remove(src)
	if err != nil {
		fmt.Println("Error removing program:", err)
		return
	}

	fmt.Println("Program removed successfully")

}

func Gobin() string {
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		// Fallback to GOPATH/bin if GOBIN is not set
		gopath := build.Default.GOPATH
		gobin = filepath.Join(gopath, "bin")
	}
	return gobin
}

func Help() {
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
}

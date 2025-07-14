# GoForge 🚀

**GoForge** is a lightweight and efficient build tool designed to simplify the development workflow for **Go (Golang)** projects. It automates repetitive tasks like project initialization, dependency management, optimization, and even cross-compilation — so you can focus on writing code, not boilerplate.

---

## ✨ Features

- 🔧 **Project Initialization** – Scaffold Go projects with sensible defaults.
- 📦 **Automatic Module Management** – No need to manually run `go mod tidy` after every change.
- 🚀 **Optimized Builds** – Easy-to-apply build flags for performance.
- 🌍 **Cross Compilation** – Build for multiple platforms with a single command.
- ⚡ **One-liner Build & Run** – Quickly test your binaries with minimal effort.

---

## 📥 Installation

Install GoForge using:

`````bash
go install github.com/ACPSOFTWARES/goforge
`````

🚀 Usage
Use goforge in your terminal to manage Go project builds and automation:

````bash
goforge <command> [args]
````
`````
| Command                          | Description                                                                |
| -------------------------------- | -------------------------------------------------------------------------- |
| `goforge new <pkg-name>`         | Create a new Go project in the current directory and initialize `go.mod`.  |
| `goforge build`                  | Build the Go project in the current directory.                             |
| `goforge run`                    | Run the compiled binary. Requires that `goforge build` has been run first. |
| `goforge build run`              | Build and immediately run the project binary.                              |
| `goforge install`                | Install the binary to `$GOBIN`. *(Currently experimental)*                 |

`````

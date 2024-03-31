Creating and running a simple Go project involves a few straightforward steps. Here's a guide to get you started:

### Step 1: Install Go

Ensure that Go is installed on your system. You can download it from the [official Go website](https://golang.org/dl/). Follow the installation instructions for your operating system.

After installation, you can verify it by opening a terminal or command prompt and running:

```sh
go version
```

This command should print the installed version of Go.

### Step 2: Set Up Your Go Workspace

As of Go 1.11, you can use Go Modules, which allow you to work outside of the traditional GOPATH-based workspace. This makes dependency management and project setup easier.

### Step 3: Create Your Project Directory

Choose or create a directory where you want your Go project to live. For this example, let's call it `hello`.

```sh
mkdir hello
cd hello
```

### Step 4: Initialize a New Module

Initialize a new module by running `go mod init` followed by the name of your module (typically the repository path, but for local projects, it can be anything).

```sh
go mod init hello
```

This command creates a `go.mod` file in your project directory, which will track your project's dependencies.

### Step 5: Write Your Go Code

Create a new file named `main.go` inside your project directory. Open it in a text editor and add the following code:

**main.go**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

This simple program will print "Hello, Go!" to the console.

### Step 6: Run Your Program

With your `main.go` file created, you can now run your program. Make sure you're in the project directory (`hello` in this case) and execute:

```sh
go run main.go
```

This command compiles and runs your Go program. You should see "Hello, Go!" printed to the console.

### Step 7: Build Your Program (Optional)

If you want to compile your program into an executable binary, use the `go build` command. This will generate an executable file in your project directory.

```sh
go build
```

After running this command, you'll find an executable named `hello` (or `hello.exe` on Windows) in your directory. You can run this executable directly:

```sh
./hello
```

Or on Windows:

```cmd
hello.exe
```

### Summary

You've now created a simple Go project, written a basic program, and learned how to run and optionally build it. Go's toolchain makes it easy to compile, run, and manage dependencies, allowing you to focus on writing your code.
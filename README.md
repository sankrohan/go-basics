Coding Exercise #1

Go to the Go Playground and starting from this empty file create the basic structure of a typical Go Application.
Define one or more variables and print their value.
Run the program without errors.


Coding Exercise #2

In VSCode in a directory called exercise2 under $GOPATH/src/master_go_programming create a file named main.go with the typical structure of any Go Application (exercise #1).
Open the terminal in VSCode or independent from VSCode. Run the program in the terminal by typing: go run main.go


Exercise #3

Compile the program from Exercise #2 manually using the go command. The executable should appear in the current working directory. Notice its name and run it!


Exercise #4

Compile the program from Exercise #2 manually using the go command. Produce an executable called my_first_go_app.exe (if you are on Windows) or my_first_go_app (if you are on Linux/Mac) in the current working directory.


Exercise #5

If you are on Windows compile the program from Exercise #2 for Linux using the go command.
If you are on Mac compile the program from Exercise #2 for Windows using the go command.
If you are on Linux compile the program from Exercise #2 for Mac using the go command.


Exercise #6

Download the following Go source code locally on your hard disk.

package main
import "fmt"
func main(){	SayHello()
}
func SayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int=10
	_=x}

Using the go command and the appropriate option, format the source code file in the idiomatic Go style.

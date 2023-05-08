//1

package main

import "fmt"

func main() {
	var fname = "Rohan"
	var lname = "Sank"

	fmt.Println("My name is", fname, lname)

	//short hand variable
	age := 10
	fmt.Print("My age is ", age)

}

//2
// go run main.go

//3
// go build main.go

//4
// GOOS=windows GOARCH=amd64 go build -o app.exe

// GOOS=linux GOARCH=amd64 go build -o linuxapp

// GOOS=darwin GOARCH=amd64 go build -o macapp

//5

// GOOS=linux GOARCH=amd64 go build -o linuxapp

// GOOS=windows GOARCH=amd64 go build -o app.exe

// GOOS=darwin GOARCH=amd64 go build -o macapp

//6

// package main

// import "fmt"

// func main() {
// 	SayHello()
// }
// func SayHello() {
// 	fmt.Println("Hello package v1.0.0!")

// 	var x int = 10
// 	_ = x
// }

//7

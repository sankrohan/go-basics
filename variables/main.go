package main

import "fmt"

// Variables

// keyword identitifer type = value
// keyword identifier = value
// keyword := value

var age int = 10
var legalAge = 18

// cannot change the value of constant
const secondsInHour = 3600

func main() {

	// short declaration operator for variables -> this is a block scope btw, it wont work outside the function
	s := "Hello World"
	fmt.Println(s)

	// muting the unused variable error by below line
	//_ = age
	fmt.Print("The sus age is", age)

	drinkinAge := 21
	fmt.Println("The drinking Age here is:", drinkinAge)

	// you can mute a global variable by using _ = variableName
	_ = legalAge
	_ = secondsInHour

	// you cannont override the short declared variable
	// s := "Hello" // this will give error
	// To override variables we use normal declaration
	// like below
	// name = "John"
	// fmt.Println(name)
	// name = "Wick"

	var name = "John"
	name = "Wick"
	fmt.Println(name)

	// Declaring multiple variables

	book, cost := "ABCD", 100
	//fmt.Println(book, cost)

	// you cannot change the short declared variable
	// but in multiple declration you can
	// because the rule for short hand is either create a new one or ATLEAST one of the variable should be new
	// example

	book, price := "XYZ", 200
	fmt.Println(book, price)

	_ = cost

	// so the variable book is now changed to XYZ

	// declaring normal variables

	var (
		salary   float64
		fname    string
		isActive bool
	)

	fmt.Println(salary, fname, isActive)

	// this return 0, "", false

	// if you want to declare variables with same type
	var a, b, c int
	fmt.Println(a, b, c)

	// this returns 0, 0, 0

	// declaring normal variables with values

	var i, j int
	i, j = 5, 8

	// swapping variables
	i, j = j, i
	fmt.Println(i, j)

	// we can use short hand declaration operator to swap variables
	x, y := 10, 5
	x, y = y, x
	fmt.Println(x, y)

	// we can also use expressions with short hand declaration

	sum := 5 + 5.55
	fmt.Println(sum)

}

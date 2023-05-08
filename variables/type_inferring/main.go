package main

import "fmt"

func main() {
	var a = 5
	var b = 5.2
	fmt.Println(a, b)

	// this will work, but if you try to swap variables of different type
	// it will throw an error

	// a, b = b, a
	// fmt.Println(a, b)
	// this will throw an error

	// so to convert variable type

	a = int(b)
	fmt.Println(a, b)

	/* zero value mecahnism
	In go, there is no such thing as uninitialized variable
	the variable will always hold a well defined value of that tyoe

	ZERO VALUES FOR EACH TYPE

	numeric types -> 0
	boolean -> false
	string -> ""(empty string)
	pointers, functions, interfaces, slices, channels, maps -> nil

	example
	var num int
	var f float64
	var str string
	var boolean bool
	fmt.Println(num, f, str, boolean)

	this will print 0 0 "" false

	*/

}

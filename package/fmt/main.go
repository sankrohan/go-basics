package main

import "fmt"

// the package fmt is used for formatting the output

func main() {
	name := "John"
	fmt.Println("Hello Mr.", name, "Wick")

	figure := "circle"
	radius := 5
	pi := 3.141

	fmt.Printf("The area of %s is %f\n", figure, pi*float64(radius*radius))

	// so % is a place holder for the value
	// %s is for string
	// %d is for integer and decimal
	// %f is for float
	// %t is for boolean
	// %v is for any value
	// %q is for quoted string
	// %T is for type of the value
	// %p is for pointer
	// %b is for binary
	// %c is for character
	// %x is for hexadecimal
	// %e is for scientific notation
	// %E is for scientific notation
	// %U is for unicode

	fmt.Printf("What is the round shape thing called, Oh its called a %q, and the type of the value is %T\n", figure, figure)
	fmt.Printf("The type of pi is %T\n", pi)

	fmt.Printf("The number 55 in binary is %08b\n", 55)

	fmt.Printf("The number 55 in hexadecimal is %x\n", 55)

	// so the 0 in %08b is for padding with 0 and 8 is for the length of the string useful for IP addresses

	// similarly to print only till 2 or 3 decimals in float we use %.2f or %.3f
	x, y := 6.6, 8.11

	fmt.Printf("x * y is %.4f\n", x*y)
	// x * y is 53.5260
	fmt.Printf("x * y is %.2f\n", x*y)
	// x * y is 53.53

}

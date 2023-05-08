package main

import "fmt"

// Constants
// constants are immutable
// constants belong to compile time and not run time
// meaning, you cannot change the value of constant during runtime and they can only be assigned during compile time
// constants can throw warn errors prior to runtime if used

// It is okay to declare a constant and not use it, but it is not okay to declare a variable and not use it
// because variables are allocated memory and constants are not

const secondsInHour = 3600
const hoursInDay = 24

// you can also declare multiple constants in one line

const (
	a = 5
	b
	c
)

// notice something strange here, the value of b and c is also 5, cus the value of a is 5 and the value of b and c is not assigned
// and in grouped constants the value of the previous constant is assigned to the next one

func main() {
	fmt.Println("seconds in hour:", secondsInHour)
	fmt.Println("hours in day:", hoursInDay)
}

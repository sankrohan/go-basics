//                                  Constants
//Sequence -> Package -> Import ->  Variables
//									Functions  (etc, the sequence after import does not matter)

package main

import "fmt"

func main() {
	fmt.Println("Hello Friend!")
	distance := 60.8                                             //distance in km
	fmt.Printf("The distance in miles is %f \n", distance*0.621) //%f is a placeholder for float value
}

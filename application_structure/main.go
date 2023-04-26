//                                  Constants
//Sequence -> Package -> Import ->  Variables
//									Functions  etc.

package main

import "fmt"

const _secondsInHour = 3600

func main() {
	fmt.Println("Hello Friend!")
	distance := 60.8 //distance in km
	fmt.Printf("The distance in miles is %f \n", distance*0.621)
}

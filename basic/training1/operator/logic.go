package main

import "fmt"

func main() {
	var right = true
	var left = false

	var rightAndLeft = right && left
	fmt.Printf("right && left \t(%t) \n", rightAndLeft)

	var rightOrLeft = right || left
	fmt.Printf("right || left \t(%t) \n", rightOrLeft)

	var rightReserve = !right
	fmt.Printf("!right \t\t(%t) \n", rightReserve)

	var leftReserve = !left
	fmt.Printf("!left \t\t(%t) \n", leftReserve)
}

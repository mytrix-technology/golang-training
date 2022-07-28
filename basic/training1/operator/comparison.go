package main

import "fmt"

func main() {
	var calculate = (((6+9)%3)*5 - 7) / 3
	var isEqual = (calculate == -2)
	var isNotEqual = (calculate == -1)

	fmt.Printf("nilai %d (%t) \n", calculate, isEqual)
	fmt.Printf("nilai %d (%t) \n", calculate, isNotEqual)
}

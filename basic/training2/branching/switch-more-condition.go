package main

import "fmt"

func main() {
	var points = 5

	switch points {
	case 10:
		fmt.Println("perfect")
	case 7, 8, 9:
		fmt.Println("awesome")
	default:
		fmt.Println("good")
	}

}

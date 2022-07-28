package main

import "fmt"

func main() {
	var points = 10

	switch points {
	case 10:
		fmt.Println("perfect")
	case 8:
		fmt.Println("awesome")
	default:
		fmt.Println("good")
	}

}

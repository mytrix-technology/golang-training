package main

import (
	"fmt"
)

func main() {
	var points = 0

	if points > 6 {
		switch points {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Awesome")
		}
	} else {
		if points == 5 {
			fmt.Println("not bad")
		} else if points == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if points == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

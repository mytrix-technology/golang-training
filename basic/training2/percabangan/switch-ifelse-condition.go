package main

import "fmt"

func main() {
	var points = 10

	switch {
	case points == 10:
		fmt.Println("perfect")
	case (points < 10) && (points > 5):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

}

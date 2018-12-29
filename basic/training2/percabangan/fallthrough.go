package main

import "fmt"

func main() {
	var points = 10

	switch {
	case points == 10:
		fmt.Println("Perfect")
	case (points < 10) && (points > 5):
		fmt.Println("Awesome")
		fallthrough
	case (points <= 5):
		fmt.Println("You need to learn")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn")
		}
	}
}

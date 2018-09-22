package main

import "fmt"

func main() {
	var one, two, three string
	one, two, three = "Bang", "Yudz", "Tampan"

	var four, five, six string = "Humble", "Bijaksana", "Dermawan"

	seven, eight, nine := "suka naik gunung", "main futsal", "poligami"

	first, isFriday, threeOnThree, swinger := 1, true, 3.3, "Guten Morgen"

	fmt.Println(one, two, three, four, five, six, seven, eight, nine)
	fmt.Println(first, isFriday, threeOnThree, swinger)
}

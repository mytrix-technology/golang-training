package main

import (
	"fmt"
	"strings"
)

func main() {
	yourHobbies("wick", "sleeping", "eating")

	// var hobbies = []string{"sleeping", "eating"}
	// yourHobbies("wick", hobbies...)
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

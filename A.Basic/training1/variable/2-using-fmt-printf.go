package main

import "fmt"

func main() {
	var firstname string = "Bang"

	var lastname string
	lastname = "Yudz"

	fmt.Printf("Guten Morgen Bang Yudz!\n")
	fmt.Printf("Guten Morgen %s %s!\n", firstname, lastname)
	fmt.Println("Guten Morgen", firstname, lastname+"!")
}

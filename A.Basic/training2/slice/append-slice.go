package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}
	var aBuah = append(buah, "Pepaya")

	fmt.Println(buah)
	fmt.Println(aBuah)
}

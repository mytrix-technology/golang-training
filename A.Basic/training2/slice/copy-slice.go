package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}
	var aBuah = []string{"Nanas", "Pepaya"}

	var copyBuah = copy(buah, aBuah)

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(copyBuah)
}

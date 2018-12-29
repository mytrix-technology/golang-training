package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}
	var aBuah = buah[0:2]
	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	fmt.Println(aBuah)
	fmt.Println(len(aBuah))
	fmt.Println(cap(aBuah))

	var bBuah = append(aBuah, "Pepaya")

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
}

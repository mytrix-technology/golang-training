package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	var aBuah = buah[0:3]
	fmt.Println(len(aBuah))
	fmt.Println(cap(aBuah))

	var bBuah = buah[1:4]
	fmt.Println(len(bBuah))
	fmt.Println(cap(bBuah))
}

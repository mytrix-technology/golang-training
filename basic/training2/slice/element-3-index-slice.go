package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}
	var aBuah = buah[0:2]
	var bBuah = buah[0:2:2]

	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	fmt.Println(aBuah)
	fmt.Println(len(aBuah))
	fmt.Println(cap(aBuah))

	fmt.Println(bBuah)
	fmt.Println(len(bBuah))
	fmt.Println(cap(bBuah))
}

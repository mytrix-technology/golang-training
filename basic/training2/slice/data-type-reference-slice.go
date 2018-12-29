package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}

	var aBuah = buah[0:3]
	var bBuah = buah[1:4]

	var aaBuah = aBuah[1:2]
	var baBuah = bBuah[0:1]

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(baBuah)

	baBuah[0] = "Nanas"

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(baBuah)
}

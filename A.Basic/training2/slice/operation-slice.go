package main

import "fmt"

func main() {
	var buah = []string{"Mangga", "Jeruk", "Apel", "Sirsak"}

	fmt.Println(buah)
	fmt.Println(buah[0:2])
	fmt.Println(buah[0:4])
	fmt.Println(buah[0:0])
	fmt.Println(buah[4:4])
	fmt.Println(buah[:])
	fmt.Println(buah[2:])
	fmt.Println(buah[:2])
}

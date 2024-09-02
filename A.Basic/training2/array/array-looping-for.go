package main

import "fmt"

func main() {
	var buah = [4]string{"Mangga", "Jeruk", "Apel", "Sirsak"}

	for i := 0; i < len(buah); i++ {
		fmt.Printf("elemen %d : %s\n", i, buah[i])
	}
}

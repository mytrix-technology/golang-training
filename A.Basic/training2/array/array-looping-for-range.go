package main

import "fmt"

func main() {
	var buah = [4]string{"Mangga", "Jeruk", "Apel", "Sirsak"}

	for i, buah := range buah {
		fmt.Printf("elemen %d : %s\n", i, buah)
	}
}

package main

import "fmt"

func main() {
	var buah = [4]string{"Mangga", "Jeruk", "Apel", "Sirsak"}

	for _, buah := range buah {
		fmt.Printf("nama buah : %s\n", buah)
	}
}

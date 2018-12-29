package main

import "fmt"

func main() {
	var angka1 = [2][3]int{[3]int{2, 3, 2}, [3]int{1, 2, 1}}
	var angka2 = [2][3]int{{2, 3, 2}, {1, 2, 1}}

	fmt.Println("angka 1", angka1)
	fmt.Println("angka 2", angka2)
}

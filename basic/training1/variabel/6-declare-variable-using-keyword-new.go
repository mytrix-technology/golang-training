package main

import "fmt"

func main() {
	name := new(string)

	fmt.Println(name)  // 0xc04202e1c0
	fmt.Println(*name) // ""
}

package main

import "fmt"

func main() {
	var ayam map[string]int
	ayam = map[string]int{}

	ayam["jantan"] = 50
	ayam["betina"] = 40

	fmt.Println("jantan", ayam["jantan"])
	fmt.Println("cacat", ayam["cacat"])
}

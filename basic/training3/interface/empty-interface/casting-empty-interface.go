package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 3
	var number = secret.(int) * 15
	fmt.Println(secret, "multiplied by 15 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")
}

package main

import "fmt"

type Person8 struct {
	name string `tag1`
	age  int    `tag2`
}

type People = Person8

func main() {
	var p1 = Person8{"wick", 21}
	fmt.Println(p1)

	var p2 = People{"wick", 21}
	fmt.Println(Person8(p2))
	fmt.Println(People(p1))
}

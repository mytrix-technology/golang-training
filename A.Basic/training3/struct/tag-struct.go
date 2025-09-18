package main

import "fmt"

type person7 struct {
	name string `tag1`
	age  int    `tag2`
}

func main() {
	var p1 = person7{"wick", 21}
	fmt.Println(p1)
}

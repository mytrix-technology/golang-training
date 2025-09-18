package main

import "fmt"

type student2 struct {
	name  string
	grade int
}

func main() {
	var s1 = student2{name: "wick", grade: 2}

	var s2 *student2 = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
}

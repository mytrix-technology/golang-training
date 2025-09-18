package main

import "fmt"

type person6 struct {
	name string
	age  int
}

func main() {
	var student6 struct {
		person6
		grade int
	}

	student6.person6 = person6{"wick", 21}
	student6.grade = 2

	fmt.Println(student6.name)
	fmt.Println(student6.age)
	fmt.Println(student6.grade)
}

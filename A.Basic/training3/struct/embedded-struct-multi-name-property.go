package main

import "fmt"

type person1 struct {
	name string
	age  int
}

type student4 struct {
	person1
	age   int
	grade int
}

func main() {
	var s1 = student4{}
	s1.name = "wick"
	s1.age = 21         // age of student
	s1.person1.age = 22 // age of person

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person1.age)
}

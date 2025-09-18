package main

import "fmt"

type person4 struct {
	name string
	age  int
}

func main() {
	var allStudents = []person4{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}

package main

import "fmt"

type person2 struct {
	name string
	age  int
}

type student5 struct {
	person2
	age   int
	grade int
}

func main() {
	var p1 = person2{name: "wick", age: 21}
	var s1 = student5{person2: p1, grade: 2}

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("grade :", s1.grade)
}

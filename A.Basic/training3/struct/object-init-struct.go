package main

import "fmt"

type student1 struct {
	name  string
	grade int
}

func main() {
	var s1 = student1{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student1{"ethan", 2}

	var s3 = student1{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
}

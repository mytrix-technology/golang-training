package main

import "fmt"

type person3 struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person3
		grade int
	}{}
	s1.person3 = person3{"wick", 21}
	s1.grade = 2

	fmt.Println("name  :", s1.person3.name)
	fmt.Println("age   :", s1.person3.age)
	fmt.Println("grade :", s1.grade)

	var s2 = struct {
		person3
		grade int
	}{
		person3: person3{"wick", 21},
		grade:   2,
	}

	fmt.Println("name  :", s2.person3.name)
	fmt.Println("age   :", s2.person3.age)
	fmt.Println("grade :", s2.grade)
}

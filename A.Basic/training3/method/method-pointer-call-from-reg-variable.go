package main

import "fmt"

func main() {
	var s1 = student2{"john wick", 21}
	s1.sayHello()

	var s2 = &student2{"ethan hunt", 22}
	s2.sayHello()
}

type student2 struct {
	name  string
	grade int
}

func (s *student2) sayHello() {
	fmt.Println("halo", s.name)
}

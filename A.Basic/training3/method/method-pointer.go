package main

import "fmt"

func main() {
	var s1 = student1{"john wick", 21}
	fmt.Println("s1 before", s1.name)
	// john wick

	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name)
	// john wick

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name)
	// ethan hunt
}

type student1 struct {
	name  string
	grade int
}

func (s student1) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student1) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

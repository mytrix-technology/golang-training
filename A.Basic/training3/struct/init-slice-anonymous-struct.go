package main

import "fmt"

type person5 struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct {
		person5
		grade int
	}{
		{person5: person5{"wick", 21}, grade: 2},
		{person5: person5{"ethan", 22}, grade: 3},
		{person5: person5{"bond", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}

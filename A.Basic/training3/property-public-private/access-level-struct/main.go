package main

import (
	"access-level-struct/library"
	"fmt"
)

func main() {
	var s1 = library.Student{Name: "ethan", Grade: 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}

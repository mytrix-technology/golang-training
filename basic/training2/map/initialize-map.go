package main

import "fmt"

func main() {
	var pitik1 = map[string]int{"jantan": 50, "betina": 40}

	var pitik2 = map[string]int{
		"jantan": 50,
		"betina": 40,
	}

	var pitik3 = map[string]int{}
	var pitik4 = make(map[string]int)
	var pitik5 = *new(map[string]int)

	fmt.Println(pitik1)
	fmt.Println(pitik2)
	fmt.Println(pitik3)
	fmt.Println(pitik4)
	fmt.Println(pitik5)
}

package main

import "fmt"

func main() {
	var pitik = map[string]int{"jantan": 50, "betina": 40}

	fmt.Println(len(pitik))
	fmt.Println(pitik)

	delete(pitik, "jantan")

	fmt.Println(len(pitik))
	fmt.Println(pitik)
}

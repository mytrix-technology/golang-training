package main

import "fmt"

func main() {
	var pitik = map[string]int{"jantan": 50, "betina": 40}
	var value, isExist = pitik["cacat"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}

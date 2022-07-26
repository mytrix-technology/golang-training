package main

import "fmt"

func main() {
	var pitiks = []map[string]string{
		{"name": "pitik biru", "gender": "jantan"},
		{"name": "pitik abang", "gender": "jantan"},
		{"name": "pitik kuning", "gender": "betina"},
	}

	for _, pitik := range pitiks {
		fmt.Println(pitik["name"], pitik["gender"])
	}
}

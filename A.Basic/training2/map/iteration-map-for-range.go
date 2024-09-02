package main

import "fmt"

func main() {
	var pitik = map[string]int{
		"jantan": 50,
		"betina": 40,
		"cacat":  34,
		"normal": 67,
	}

	for key, val := range pitik {
		fmt.Println(key, "  \t:", val)
	}
}

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		for j := i; j <= 10; j++ {
			fmt.Println(j, " ")
		}

		fmt.Println()
	}
}

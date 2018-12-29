package main

import (
	"fmt"
)

func main() {
outerloop:
	for i := 0; i <= 10; i++ {
		for j := i; j <= 10; j++ {
			if i == 7 {
				break outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

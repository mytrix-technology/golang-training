package main

import "fmt"

func main() {
	var points = 9

	if points == 10 {
		fmt.Printf("LULUS SEMPURNA. Nilai anda : %d \n", points)
	} else if points > 7 && points < 10 {
		fmt.Printf("LULUS. Nilai anda : %d \n", points)
	} else if points >= 5 && points < 7 {
		fmt.Printf("REMIDIAL. Nilai anda : %d \n", points)
	} else {
		fmt.Printf("TIDAK LULUS. Nilai anda : %d \n", points)
	}

}

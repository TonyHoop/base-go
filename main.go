package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := 1 + rand.Intn(10)

	fmt.Println("Попробуй угадать число")

	for {
		var answer int
		fmt.Scan(&answer)

		if answer < number {
			fmt.Println("Больше")
		} else if answer > number {
			fmt.Println("Меньше")
		} else {
			fmt.Println("Угадал")
			break
		}
	}
}

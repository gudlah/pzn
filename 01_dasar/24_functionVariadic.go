package main

import "fmt"

func sumAll(name string, numbers ...int) (int, string) {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total, name
}

func functionVariadic() {
	total, _ := sumAll("Bagus", 11, 12, 13, 14)
	fmt.Println(total)

	slice := []int{10, 20, 30, 40, 50}
	total, _ = sumAll("Bagus", slice...)
	fmt.Println(total)
}

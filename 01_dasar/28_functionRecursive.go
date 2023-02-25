package main

import "fmt"

func faktorialLoops(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func faktorialRekursif(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRekursif(value-1)
	}
}

func functionRecursive() {
	loop := faktorialLoops(10)
	fmt.Println(loop)

	rekursif := faktorialRekursif(10)
	fmt.Println(rekursif)
}

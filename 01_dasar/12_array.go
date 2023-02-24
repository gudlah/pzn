package main

import "fmt"

func array() {
	var names [3]string

	names[0] = "Bagus"
	names[1] = "Anjar"
	names[2] = "Sadewa"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		70,
		80,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(lagi)
	fmt.Println(len(lagi))
}

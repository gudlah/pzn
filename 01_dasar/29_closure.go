package main

import "fmt"

func main() {
	name := "Bagus"
	counter := 0

	increament := func() {
		name := "Anjar"
		fmt.Println("Increament")
		counter++
		fmt.Println(name)
	}

	increament()
	increament()

	fmt.Println(counter)
	fmt.Println(name)
}

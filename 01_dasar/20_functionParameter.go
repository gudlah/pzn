package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Halo", firstName, lastName)
}

func functionParameter() {
	firstName := "Bagus"
	sayHelloTo(firstName, "Anjar Sadewa")
}

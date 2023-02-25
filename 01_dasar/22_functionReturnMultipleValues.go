package main

import "fmt"

func getBio() (string, int, string) {
	return "Bagus", 1998, "Anjar Sadewa"
}

func functionReturnMultipleValues() {
	firstName, _, lastName := getBio()
	fmt.Println("Nama : " + firstName + " " + lastName)
	// fmt.Println("Umur :", age)
}

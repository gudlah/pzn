package main

import "fmt"

func constant() {
	const (
		firstName string = "Bagus"
		lastName         = "Anjar Sadewa"
		value            = 1000
	)

	// fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// Bakal Error
	// firstName = "Tidak Bisa Diubah"
	// lastName = "Tidak Bisa Diubah"
	// value = 4000
}

package main

import (
	"fmt"
)

func variable() {
	var name string

	name = "Bagus Anjar"
	fmt.Println(name)

	name = "Bagus Sadewa"
	fmt.Println(name)

	var friendName = "Joko Satrio"
	fmt.Println(friendName)

	var age = 24
	fmt.Println(age)

	province := "Jawa Timur"
	fmt.Println(province)

	province = "Jawa Barat"
	fmt.Println(province)

	var (
		firstName = "Bagus"
		lastName  = "Waskito Aji Saputro"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

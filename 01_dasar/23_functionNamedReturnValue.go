package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Bagus"
	lastName = "Sadewa"
	middleName = "Anjar"
	return firstName, lastName, middleName
}

func functionNamedReturnFunction() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	a, b, c := getBiodata()
	fmt.Println("Nama :", a)
	fmt.Println("Umur :", b)
	fmt.Println("Alamat :", c)
}

func getBiodata() (name string, age int, address string) {
	name = "Bagus Anjar Sadewa"
	age = 24
	address = "Bekasi"
	return
}

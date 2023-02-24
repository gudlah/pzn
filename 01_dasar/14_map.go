package main

import "fmt"

func map14() {
	person := map[string]string{
		"name":    "Bagus",
		"address": "Bekasi",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Bagus"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}

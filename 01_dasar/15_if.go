package main

import "fmt"

func if15() {
	var name = "Sadewa"

	if name == "Bagus" {
		fmt.Println("Halo Bagus")
	} else if name == "Anjar" {
		fmt.Println("Halo Anjar")
	} else if name == "Sadewa" {
		fmt.Println("Halo Sadewa")
	} else {
		fmt.Println("Hi, kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}

package main

import "fmt"

func switch16() {
	name := "Sadewa"

	switch name {
	case "Bagus":
		fmt.Println("Halo Bagus")
	case "Anjar":
		fmt.Println("Halo Anjar")
	case "Sadewa":
		fmt.Println("Halo Sadewa")
	default:
		fmt.Println("Hi, kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}

package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Halo Bro"
	} else {
		return "Halo " + name
	}
}

func functionReturnValue() {
	name := getHello("Bagus")
	fmt.Println(name)
	fmt.Println(getHello(""))
}

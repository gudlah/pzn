package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func functionAsValue() {
	goodBye := getGoodBye
	result := goodBye("Bagus")
	fmt.Println(result)
}

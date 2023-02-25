package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", filter(nameFiltered))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func functionAsParameter() {
	sayHelloWithFilter("Apa", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

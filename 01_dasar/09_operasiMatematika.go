package main

import "fmt"

func operasiMatematika() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// Augmented Assigment
	var i = 10
	i += 10
	fmt.Println(i)

	// Unary Operator
	i++
	var negative = -100
	var positive = +100
	fmt.Println(i)
	fmt.Println(negative)
	fmt.Println(positive)
}

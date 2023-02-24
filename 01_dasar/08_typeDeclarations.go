package main

import (
	"fmt"
)

func typeDeclarations() {
	type NoKTP string
	type Married bool

	var ktpKu NoKTP = "11111"
	var marriedStatus Married = true
	fmt.Println(ktpKu)
	fmt.Println(marriedStatus)
}

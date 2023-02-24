package main

import "fmt"

func main() {
	// var counter = 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke", counter)
	}

	names := []string{"Bagus", "Anjar", "Sadewa"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("Index", i, "=", name)
	}

	person := make(map[string]string)
	person["name"] = "Bagus"
	person["address"] = "Bekasi"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

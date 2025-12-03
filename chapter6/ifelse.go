package main

import "fmt"

func main() {
	// age := 12
	// if age >= 18 {
	// 	fmt.Println("Adult Person")
	// } else {
	// 	fmt.Println("Not adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Adult Person")
	// } else if age >= 12 {
	// 	fmt.Println("Teenager")
	// } else {
	// 	fmt.Println("child")
	// }

	// var role = "admin"
	// var hasPermisssions = false

	// if role == "admin" || hasPermisssions {
	// 	fmt.Println("yes")
	// }

	// we can declare a variable inside if construct
	if age := 16; age >= 18 {
		fmt.Println("Adult Person", age)
	} else if age >= 12 {
		fmt.Println("teenage", age)
	}

	// go doesnot have ternary operator
}

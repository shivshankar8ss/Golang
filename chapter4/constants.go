package main

import "fmt"

const age = 30

func main() {
	const name string = "shivshankar"
	// name = "golang"
	fmt.Println(age)

	// multiple constants : grouping them
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}

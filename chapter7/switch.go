package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 1

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default: //optional
	// 	fmt.Println("Other")
	// }

	// multiple condition switch
	println(time.Now().Date())
	println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("workday")
	}
}

package main

import "fmt"

// only for exists in go for looping
func main() {
	// while loop
	// i:=1
	// for i<=5{
	// 	fmt.Println(i)
	// 	i+=1
	// }

	// infinite loop
	// for {
	// 	println("shiv")
	// }

	// classic for loop
	// for i := 0; i <= 6; i++ {
	// 	// break
	// 	// continue
	// 	fmt.Println(i)
	// }

	for i:= range 10{
		fmt.Println(i)
	}
}

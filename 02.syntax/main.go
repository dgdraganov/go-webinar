package main

import "fmt"

func main() {
	var number int = 70

	if number%2 == 0 {
		fmt.Printf("Number %d is even.\n", number)
		fmt.Printf("Variable type is %T.\n", number)
		fmt.Printf("Number in binary %b.\n", number)
		fmt.Printf("Number in hex %x.\n", number)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	condition := true

	for condition {
		condition = false
	}

	// unused := 55 // ERROR
}

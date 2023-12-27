package main

import (
	"fmt"
)

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

	cond, err := IsAdult(15)
	if err != nil {
		fmt.Printf("check if adult: %s\n", err)
	}

	fmt.Printf("is adult: %t\n", cond)
}

func IsAdult(age int) (bool, error) {
	if age < 0 {
		return false, fmt.Errorf("age cannot be less than 0, age: %d", age)
	}
	return age >= 18, nil
}

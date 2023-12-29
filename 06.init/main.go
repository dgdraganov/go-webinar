package main

import "fmt"

func init() {
	fmt.Println("this will be executed on package initialization")
}

func main() {
	fmt.Println("Hello world!")
}

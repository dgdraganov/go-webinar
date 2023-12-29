package main

import "fmt"

type Status string

const (
	Success    Status = "success"
	Failed     Status = "failed"
	Processing Status = "processing"
)

type Weekday int

const (
	Monday    Weekday = iota // 0
	Tuesday                  // 1
	Wednesday                // 2
	Thursday                 // 3
	Friday                   // 4
	Saturday                 // 5
	Sunday                   // 6
)

func main() {
	var status Status = Failed
	fmt.Println(status)

	var day Weekday = Friday
	fmt.Println(day)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("西华，今天是周几？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	fmt.Println(today, "西华")
}

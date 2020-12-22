package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "西"
	array[1] = "华"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
}

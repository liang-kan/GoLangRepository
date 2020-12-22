package main

import "fmt"

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	var s []int = primes[1:4]
	var ss []int = primes[1:2]
	var sss []int = primes[1:3]
	var ssss []int = primes[1:6]
	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(sss)
	fmt.Println(ssss)

	var x []int = primes[0:1]
	var xx []int = primes[1:1]
	fmt.Println(x)
	fmt.Println(xx)
}

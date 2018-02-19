package main

// Project Euler Problem #7
// ------------------------
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n <= 3 {
		return n == 2 || n == 3
	}

	sqrt := math.Sqrt(float64(n))

	for i := 5; i < int(sqrt)+1; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	counter := 1
	num := 1

	for counter < 10001 {
		num += 2
		if IsPrime(num) {
			counter++
		}
	}
	fmt.Println(num)
}

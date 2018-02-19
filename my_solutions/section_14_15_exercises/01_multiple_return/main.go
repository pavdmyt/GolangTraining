package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

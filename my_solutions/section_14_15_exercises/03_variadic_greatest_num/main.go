package main

import "fmt"

func greatestNum(sn ...int) int {
	var g int
	for _, i := range sn {
		if i > g {
			g = i
		}
	}
	return g
}

func main() {
	items := []int{2, 3, 11, 15, 3, 5, 19, 0, 1}
	fmt.Println(greatestNum(items...))
}

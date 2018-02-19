package main

import "fmt"

func foo(sn ...int) {
	for i, n := range sn {
		if i == len(sn)-1 {
			fmt.Println(n)
		} else {
			fmt.Printf("%d ", n)
		}
	}
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)

	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

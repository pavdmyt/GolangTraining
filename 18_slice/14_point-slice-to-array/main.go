package main

import "fmt"

func main() {
	var arr [12]int
	fmt.Printf("%v -- %T \n", arr, arr)

	s := arr[:]
	fmt.Printf("%v -- %T \n", s, s)

	for i := 33; i < 33+len(arr); i++ {
		s[i-33] = i
	}
	fmt.Printf("arr: %v \n", arr)
	fmt.Printf("s: %v \n", s)
}

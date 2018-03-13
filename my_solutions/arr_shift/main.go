package main

import "fmt"

func shift(arr []int, shift int) []int {
	l := len(arr)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = arr[(i+shift+1)%l]
	}
	return res
}

func main() {
	inArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	outArr := shift(inArr, 4)
	fmt.Println(outArr)
}

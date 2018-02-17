package main

import "fmt"

func Shift(arr []int, shift int) []int {
	l := len(arr)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = arr[(i+shift+1)%l]
	}
	return res
}

func main() {
	in_arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	out_arr := Shift(in_arr, 4)
	fmt.Println(out_arr)
}

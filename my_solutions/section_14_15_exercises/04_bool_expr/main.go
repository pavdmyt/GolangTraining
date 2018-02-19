package main

import "fmt"

func main() {
	res := (true && false) || (false && true) || !(false && false)
	fmt.Println(res)
}

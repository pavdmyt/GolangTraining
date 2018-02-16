package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("> enter number: ")
	fmt.Scan(&num1)

	fmt.Print("> enter larger number: ")
	fmt.Scan(&num2)

	if num2 < num1 {
		fmt.Printf("Error: %d < %d \n", num2, num1)
	} else {
		fmt.Printf("Reminder: %d \n", num2%num1)
	}
}

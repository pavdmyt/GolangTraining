package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	qty := 100
	wg.Add(qty)
	work := gen(qty)
	res := factorial(work, qty)

	for n := range res {
		fmt.Println(n)
	}
}

func gen(qty int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < qty; i++ {
			n := rand.Intn(11)
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int, wNum int) <-chan int {
	out := make(chan int)

	for n := range in {
		n := n
		go func() {
			out <- fact(n)
			wg.Done()
		}()
	}

	// Waiting and closing should be
	// in a separate goroutine!!!
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1
	for i := 2; i <= n; i++ {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/

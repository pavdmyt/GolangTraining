package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	qty := 5000 // scope of work
	workers := 1000

	dataCh := gen(qty)
	resCh := factFanOut(dataCh, workers)

	for res := range resCh {
		fmt.Println(res)
	}
}

func gen(qty int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < qty; i++ {
			n := rand.Intn(15)
			out <- n
		}
		close(out)
	}()
	return out
}

func factFanOut(in <-chan int, wNum int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for w := 0; w < wNum; w++ {
		wg.Add(1)
		go func() {
			for n := range in {
				out <- fact(n)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

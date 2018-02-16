package main

func main() {
	var sum int
	for i := 3; i < 1e8; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	println(sum)
}

/*
$ time go run main.go
2333333316666668
go run main.go  0.34s user 0.13s system 63% cpu 0.741 total
*/

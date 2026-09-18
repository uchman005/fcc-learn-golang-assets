package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	intChan := make(chan int)
	go fibonacci(n, intChan)

	for data := range intChan {
		fmt.Println(data)
	}
}

// TEST SUITE - Don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	defer close(ch)
	x, y := 0, 1
	for range n {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
}

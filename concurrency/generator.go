package main

import "fmt"

func main() {
	fmt.Println("No bottles of beer on the wall")

	for i := range Count(1, 99) {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
	}

	fmt.Println(100, "bottles of beer on the wall")
}

func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

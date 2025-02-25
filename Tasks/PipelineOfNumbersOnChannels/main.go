package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()
	return out
}

func main() {
	for ch := range sq(gen(7, 5, 3, 2, 9)) {
		fmt.Println(ch)
	}
}

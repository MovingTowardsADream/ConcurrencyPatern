package main

import (
	"fmt"
	"sync"
)

func mergeChan(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	for _, c := range cs {
		go func() {
			for i := range c {
				out <- i
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

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	ch3 := mergeChan(ch1, ch2)
	for i := range ch3 {
		fmt.Println(i)
	}
}

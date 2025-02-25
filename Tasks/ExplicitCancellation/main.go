package main

import (
	"fmt"
	"sync"
	"time"
)

func mergeChannelsWithsDone(done chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))

	go func() {
		for _, ch := range cs {
			go func() {
				defer wg.Done()
				for c := range ch {
					select {
					case out <- c:
					case <-done:
						return
					}
				}
				fmt.Println("I'm stopping now")
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	first, second, third := make(chan int), make(chan int), make(chan int)

	go func() {
		first <- 1
		time.Sleep(2 * time.Second)
		first <- 2
		close(first)
	}()

	go func() {
		second <- 3
		time.Sleep(2 * time.Second)
		second <- 4
		close(second)
	}()

	go func() {
		third <- 5
		time.Sleep(2 * time.Second)
		third <- 6
		close(third)
	}()

	done := make(chan struct{})
	in := mergeChannelsWithsDone(done, first, second, third)

	go func() {
		defer close(done)
		for {
			select {
			case res, ok := <-in:
				if !ok {
					return
				}
				fmt.Println(res)
			case <-time.After(1 * time.Second):
				return
			}
		}
	}()
	<-done
	fmt.Println("End")
}

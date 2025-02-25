package main

import (
	"fmt"
	"sync"
	"time"
)

const amountGorutine = 2

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	ch := make(chan struct{}, amountGorutine)
	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-ch }()
			// Main work
			time.Sleep(1 * time.Second)
			fmt.Println("Yes")
			//
		}()
	}
	wg.Wait()
}

package main

import (
	"context"
	"fmt"
	"time"
)

func longFunc() int {
	time.Sleep(2 * time.Second)
	return 0
}

func normalTimeFunc() {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var value int

	ch := make(chan int)

	go func() {
		value = longFunc()
		close(ch)
	}()

	select {
	case <-ctxWithTimeout.Done():
	case _, _ = <-ch:
		fmt.Println(value)
	}

}

func main() {
	normalTimeFunc()
}

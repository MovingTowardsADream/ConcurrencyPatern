package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	mu   sync.Mutex
	head *Node
}

func (s *Stack) Push(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tmp := s.head
	s.head = &Node{value, tmp}
}

func (s *Stack) Top() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.head != nil {
		return s.head.Value
	}
	return -1
}

func (s *Stack) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.head != nil {
		s.head = s.head.Next
	}
}

func main() {
	stack := Stack{head: nil}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				stack.Push(i)
			}
		}()
	}
	wg.Wait()
	for i := 0; i < 100; i++ {
		fmt.Println(stack.Top())
		stack.Pop()
	}
}

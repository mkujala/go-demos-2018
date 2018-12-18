// Stack FIFO (queue)
package main

import (
	"fmt"
	"os"
)

type stack []int

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(99)
	s.push(88)
	s.push(3)
	fmt.Println(s)
	fmt.Println("pop:", s.pop())
	fmt.Println("pop:", s.pop())
	fmt.Println("pop:", s.pop())
	fmt.Println(s)
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() int {
	var popped int
	if len((*s)) > 0 {
		popped = (*s)[0]
		*s = (*s)[1:]
	} else {
		fmt.Println("Stack is empty!")
		os.Exit(1)
	}
	return popped
}

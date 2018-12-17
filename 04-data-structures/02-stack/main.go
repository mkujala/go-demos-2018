package main

import "fmt"

type element struct {
	data interface{}
	next *element
}

type stack struct {
	head *element
	Size int
}

func main() {
	stk := new(stack)
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	stk.Push(10)
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	stk.Push(12)
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	stk.Push(8)
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	stk.Push(6)
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	fmt.Println("\npop item:", stk.Pop())
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	fmt.Println("pop item:", stk.Pop())
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	fmt.Println("pop item:", stk.Pop())
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	fmt.Println("pop item:", stk.Pop())
	fmt.Println("head:", stk.head, "Size:", stk.Size)

	fmt.Println("pop item:", stk.Pop())
}

func (stk *stack) Push(data interface{}) {
	element := new(element)
	element.data = data
	element.next = stk.head
	stk.head = element
	stk.Size++
}

func (stk *stack) Pop() interface{} {
	if stk.head == nil {
		return "Stack is empty!"
	}
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	return r
}

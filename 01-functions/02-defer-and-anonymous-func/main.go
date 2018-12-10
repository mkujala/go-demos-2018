package main

import (
	"fmt"
)

func main() {
	defer1()
	line()

	defer2()
	line()

	defer3()
	line()

	// anonymous function called with parameter 3
	func(i int) { fmt.Println(i) }(3)
}

func defer1() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // 2, 1, 0
	}
}

func defer2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }() // 3, 3, 3
	}
}

func defer3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Println(n) }(i) // 2, 1, 0
	}
}

func line() {
	fmt.Println("-------------------------")
}

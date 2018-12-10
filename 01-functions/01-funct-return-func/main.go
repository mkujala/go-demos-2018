package main

import (
	"fmt"
)

func main() {
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result11 := addFive(6)
	result20 := addFive(15)

	fmt.Printf("%v\n%v\n", result11, result20)
}

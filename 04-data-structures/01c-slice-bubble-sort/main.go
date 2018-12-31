package main

import (
	"fmt"
)

type arr = []int

func main() {
	var slice = arr{9, 1, 2, 767, 123, 9, 1337, 9, 23, 4123, 78, 243}
	var slice2 = arr{1, 2, 9, 3}
	fmt.Print("\n------------------------\n")
	fmt.Println("unsorted slice1:", slice)
	bubbleSort(&slice)
	fmt.Println("sorted slice1:", slice)
	fmt.Print("\n------------------------\n")

	fmt.Println("unsorted slice2:", slice2)
	bubbleSort(&slice2)
	fmt.Println("sorted slice2:", slice2)
}

func bubbleSort(s *arr) {
	iterations := 0
	bubbleLength := len(*s) - 1
	for i := 0; i < bubbleLength; i++ {
		for i := 0; i < len(*s)-1; i++ {
			if (*s)[i] > (*s)[i+1] {
				(*s)[i], (*s)[i+1] = (*s)[i+1], (*s)[i]
			}
			iterations++
		}
		bubbleLength--
	}
	fmt.Println("\ntotal number of iterations:", iterations)
}

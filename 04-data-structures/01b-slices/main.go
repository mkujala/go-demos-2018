// Return first non-repeating integer in an array
package main

import "fmt"

type nonRep = (map[int]int)

func main() {
	slice1 := []int{1, 2, 1, 2, 1, 3, 4, 1, 2, 3, 4}
	slice2 := []int{9, 2, 2, 2, 2, 9, 9, 9, 2, 9, 1337}
	fmt.Println("slice1:", slice1)
	firstNonRep(slice1)

	fmt.Print("\n")

	fmt.Println("slice2:", slice2)
	firstNonRep(slice2)
}

func firstNonRep(s []int) {
	nr := nonRep{}
	for _, i := range s {
		// fmt.Println(nr)
		_, ok := nr[i]
		if ok {
			nr[i] = nr[i] + 1
		} else {
			nr[i] = 0
			if len(nr) == 3 {
				fmt.Println("first non-repeating value:", i)
				break
			}
		}
	}
}

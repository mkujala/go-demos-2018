package main

import "fmt"

type arr = []int

func main() {
	var slice = arr{38, 27, 43, 3, 14, 16, 8}
	l := 0              // left pointer
	r := len(slice) - 1 // right pointer
	fmt.Println(slice)
	// fmt.Printf("l: %v, r: %v\n", l, r)
	sort(slice, l, r)
}

func sort(a arr, l, r int) {
	fmt.Printf("splitting l:%v r:%v \n", l, r)
	if l < r {
		m := (l + r) / 2
		fmt.Printf("1) l:%v, m:%v, r:%v\n", l, m, r)

		sort(a, l, m) // continue to row 23 after m=l
		fmt.Println("-----sort L done-----")
		fmt.Printf("2) l:%v, m:%v, r:%v\n", l, m, r)
		sort(a, m+1, r) // continue to row 24 after m+1=r
		fmt.Println("-----sort R done-----")
		merge(a, l, m, r)
	}
}

func merge(a arr, l, m, r int) {
	fmt.Printf("merge l:%v m:%v r:%v \n", l, m, r)
	n1 := m - l + 1
	n2 := r - m
	leftArr := make(arr, n1)
	rightArr := make(arr, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = a[l+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = a[m+j+1]
	}

	// Merge the temp arrays

	// Initial indexes of first and second subarrays
	i := 0
	j := 0
	// Initial index of merged subarry array
	k := l
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			a[k] = leftArr[i]
			i++
		} else {
			a[k] = rightArr[j]
			j++
		}
		k++
	}

	// Copy remaining elements of L[] if any
	for i < n1 {
		a[k] = leftArr[i]
		i++
		k++
	}

	// Copy remaining elements of R[] if any
	for j < n2 {
		a[k] = rightArr[j]
		j++
		k++
	}

	fmt.Println("After merge")
	printArray(a)
}

func printArray(a arr) {
	n := len(a)
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

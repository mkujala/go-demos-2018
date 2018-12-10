package main

import (
	"fmt"
	"strconv"
)

var inputs []int
var endProgram = false
var maxInput = 9999
var minInput = -9999

func main() {
	for {
		input := getInput()
		if !endProgram {
			inputs = append(inputs, input)
		} else {
			fmt.Println(endInput())
			break
		}
	}
}

func getInput() int {
	var input string
	fmt.Printf("Give number between %v to %v, 0 ends program: ", minInput, maxInput)
	fmt.Scanln(&input)
	inputINT, err := strconv.Atoi(input)
	if err != nil || inputINT > maxInput || inputINT < minInput {
		fmt.Printf("Invalid input: %s, Exiting program.\n\n", input)
	}
	if inputINT == 0 || inputINT > maxInput || inputINT < minInput {
		endProgram = true
	}
	return inputINT
}

func endInput() string {
	min := inputs[0]
	max := inputs[0]
	for _, i := range inputs {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}
	return fmt.Sprintf("You gave following numbers: %v \nmin: %v, max: %v\n", inputs, min, max)
}

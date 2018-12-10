package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFilename = "text.txt"

type counts = (map[string]int)

func main() {
	text := string(readFile(inputFilename))
	result := countWords(text)
	printResults(result)
}

func countWords(t string) counts {
	counts := counts{}
	sText := strings.Fields(t)
	for _, word := range sText {
		_, ok := counts[word]
		if ok {
			counts[word] = counts[word] + 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func printResults(r counts) {
	for key := range r {
		fmt.Printf("%s | %d\n---------------------\n", key, r[key])
	}
}

func readFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

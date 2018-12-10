package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var inputFilename = "text.txt"

type counts = (map[string]int)

func main() {
	text := string(readFile(inputFilename))
	cleanedText := removeChars(text, "[,.!?“”]")
	words := countWords(cleanedText)
	printResults(words)
}

func countWords(t string) counts {
	counts := counts{}
	sText := strings.Fields(t) // split string to []string
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

func removeChars(text, chars string) string {
	parse, err := regexp.Compile(chars)
	if err != nil {
		log.Fatal("Error removing chars:", chars, "\n", err)
	}
	cleanedText := parse.ReplaceAllString(text, "")
	return cleanedText
}

func readFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

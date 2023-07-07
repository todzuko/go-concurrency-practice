package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	data1 := []string{
		"1",
		"10",
		"string",
		"444",
		"some text here",
		"aaaaa",
	}

	printc := make(chan string)
	processedPrintc := make(chan string)
	go stringProccessing(data1, printc)
	go numProccessor(printc, processedPrintc)
	go stringPrinter(processedPrintc)

	var input string
	fmt.Scanln(&input)
}

func stringProccessing(data []string, printc chan string) {
	for _, val := range data {
		if utf8.RuneCountInString(val) > 0 {
			printc <- val
		}
	}
}

func numProccessor(printc, proccessedPrintc chan string) {
	for val := range printc {
		if isNumber(val) {
			proccessedPrintc <- fmt.Sprintf("This is a number %s", val)
		} else if utf8.RuneCountInString(val) > 5 {
			proccessedPrintc <- val
		}
	}
	close(printc)
}

func stringPrinter(proccessedPrintc chan string) {
	for val := range proccessedPrintc {
		fmt.Println(val)
	}
	close(proccessedPrintc)
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

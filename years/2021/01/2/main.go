package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)

func main() {
	scanner, file := getInput()
	count := 0
	var rollingWindow [3]int
	for i := 0; i < 3 && scanner.Scan(); i++ {
		rollingWindow[0] = atoi(scanner.Text())
	}
	prev := sum(rollingWindow)
	for scanner.Scan() {
		shift(&rollingWindow, atoi(scanner.Text()))
		curr := sum(rollingWindow)
		if prev < curr {
			count++
		}
		prev = curr
	}
	file.Close()
	log.Printf("Count: %v", count)
}

func getInput() (*bufio.Scanner, *os.File) {
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(path.Join(path.Dir(filename), "../input.txt"))
	if err != nil {
		log.Fatal("Failed to open input file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, file
}

func atoi(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Failed to convert string to number: %v", text)
	}
	return value
}

func sum(array [3]int) int {
	result := 0
	for _, value := range array {
		result += value
	}
	return result
}

func shift(array *[3]int, value int) {
	array[0] = array[1]
	array[1] = array[2]
	array[2] = value
}

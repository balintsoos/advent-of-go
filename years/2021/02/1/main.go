package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	scanner, file := getInput()
	horizontalPos := 0
	verticalPos := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		amount := atoi(fields[1])
		switch fields[0] {
		case "forward":
			horizontalPos += amount
		case "up":
			verticalPos -= amount
		case "down":
			verticalPos += amount
		}
	}
	result := horizontalPos * verticalPos
	file.Close()
	log.Printf("Result: %v", result)
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

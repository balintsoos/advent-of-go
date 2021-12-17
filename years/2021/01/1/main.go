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
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(path.Join(path.Dir(filename), "../input.txt"))

	if err != nil {
		log.Fatal("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0
	scanner.Scan()
	prev := getValue(scanner)
	for scanner.Scan() {
		curr := getValue(scanner)
		if prev < curr {
			count += 1
		}
		prev = curr
	}

	file.Close()
	log.Printf("Count: %v", count)
}

func getValue(scanner *bufio.Scanner) int {
	text := scanner.Text()
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Failed to convert string to number: %v", text)
	}
	return value
}

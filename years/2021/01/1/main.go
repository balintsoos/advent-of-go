package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(path.Join(path.Dir(filename), "input.txt"))

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Close()
}

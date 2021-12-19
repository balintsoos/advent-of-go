package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"runtime"
)

func main() {
	scanner, file := getInput()
	var sum [12]int
	count := 0
	for scanner.Scan() {
		count++
		bit := scanner.Text()
		for i := 0; i < 12; i++ {
			sum[i] += int(bit[i] - '0')
		}
	}
	file.Close()
	var gammaRateDecimal float64 = 0
	var epsilonRateDecimal float64 = 0
	half := count / 2
	for i := 0; i < 12; i++ {
		if sum[i] > half {
			gammaRateDecimal += math.Pow(2, float64(11-i))
		} else {
			epsilonRateDecimal += math.Pow(2, float64(11-i))
		}
	}
	result := gammaRateDecimal * epsilonRateDecimal
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

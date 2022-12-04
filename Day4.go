package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Define the path of the file to read.
	filePath := "day4input.txt"

	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		// If an error occurred, print it and return.
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner for the file.
	scanner := bufio.NewScanner(file)

	counter := 0
	counter2 := 0

	for scanner.Scan() {
		x1, y1, x2, y2 := getOrderedPairs(scanner.Text())

		// Part 1 Check
		if (x1 >= x2 && y1 <= y2) || (x1 <= x2 && y1 >= y2) {
			counter++
		}

		// Part 2 Check
		if x1 <= x2 && y1 >= x2 || x2 <= x1 && y2 >= x1 {
			counter2++
		}
	}

	fmt.Println("Part 1: ", counter)
	fmt.Println("Part 2: ", counter2)

}

func getOrderedPairs(pairs string) (int, int, int, int) {
	numbers := strings.Split(pairs, ",")
	group1 := strings.Split(numbers[0], "-")
	group2 := strings.Split(numbers[1], "-")
	x1, _ := strconv.Atoi(group1[0])
	y1, _ := strconv.Atoi(group1[1])
	x2, _ := strconv.Atoi(group2[0])
	y2, _ := strconv.Atoi(group2[1])
	return x1, y1, x2, y2
}

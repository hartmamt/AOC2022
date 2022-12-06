package main

import (
	"bufio"
	"fmt"
	"os"
)

func characterExists(char string, marker []string) (bool, int) {
	for i, ch := range marker {
		if ch == char {
			return true, i
		}
	}
	return false, 0
}

func allItemsUnique(words []string) bool {
	// Create a map to store the items we have seen
	seen := make(map[string]bool)

	// Iterate over the elements in the slice
	for _, word := range words {
		// If the item is already in the map, return false
		if seen[word] {
			return false
		}

		// Add the item to the map
		seen[word] = true
	}

	// If we reach this point, all items are unique
	return true
}

func findMarker(line string, counter int) ([]string, int) {
	var markerBetter []string
	// sliding window
	for i, ch := range line {
		markerBetter = append(markerBetter[:i], string(ch))
		if i > counter-1 {
			if allItemsUnique(markerBetter[i-counter : i]) {
				return markerBetter[i-counter : i], i
			}
		}
	}
	return nil, -1
}

func main() {
	// Define the path of the file to read.
	filePath := "day6input.txt"

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

	for scanner.Scan() {

		line := scanner.Text()

		fmt.Println("Part One:")
		fmt.Println(findMarker(line, 4))
		fmt.Println("Part Two:")
		fmt.Println(findMarker(line, 14))
	}
}

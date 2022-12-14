package main

import (
	"bufio"
	"fmt"
	"os"
)

func allItemsUnique(letters []string) bool {
	// Create a map to store the items we have seen
	seen := make(map[string]bool)

	// Iterate over the elements in the slice
	for _, letter := range letters {
		if seen[letter] {
			return false
		}

		// Add the item to the map
		seen[letter] = true
	}
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

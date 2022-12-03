package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Open the file.
	file, err := os.Open("rucksack.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner for the file.
	scanner := bufio.NewScanner(file)

	partOneAnswer := 0
	badgeCount := 0
	lineNumber := 0

	// Create a slice to store the previous two items.
	prevItems := make([]string, 0)

	// Read the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		firstCompartment, secondCompartment := splitRuckSacks(line)

		firstCompartmentCounts := getCounts(firstCompartment)
		secondCompartmentCounts := getCounts(secondCompartment)

		// create a map of runes to ints
		compartments := make([]map[rune]int, 2)
		compartments[0] = firstCompartmentCounts
		compartments[1] = secondCompartmentCounts

		partOneAnswer += findItemInAllPacks(compartments)

		// Increment the line number.
		lineNumber++
		prevItems = append(prevItems, line)

		// If we have reached the third line, print it and the previous two lines.
		if lineNumber%3 == 0 {
			// create a map of runes to ints
			threeCompartments := make([]map[rune]int, 3)
			threeCompartments[0] = getCounts(prevItems[0])
			threeCompartments[1] = getCounts(prevItems[1])
			threeCompartments[2] = getCounts(prevItems[2])
			badgeCount = badgeCount + findItemInAllPacks(threeCompartments)

			// reset the slice
			prevItems = make([]string, 0)
		}
	}

	fmt.Println("The answer for part one is:", partOneAnswer)
	fmt.Println("The answer for part two is:", badgeCount)
}

func getPriority(char rune) int {
	position := unicode.ToLower(char) - rune('a') + 1
	if unicode.IsUpper(char) {
		position = position + 26
	}
	return int(position)
}

func splitRuckSacks(sack string) (string, string) {
	halfLength := len(sack) / 2
	firstCompartment := sack[:halfLength]
	secondCompartment := sack[halfLength:]
	return firstCompartment, secondCompartment
}

func getCounts(sack string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range sack {
		counts[char]++
	}
	return counts
}

func findItemInAllPacks(maps []map[rune]int) int {
	count := 0
	for item, _ := range maps[0] {
		existsInAllMaps := true

		// Iterate over the maps in the slice and check if the current item exists in all of them.
		for _, m := range maps {
			// If the current item does not exist in the current map, set the existsInAllMaps flag to false.
			if _, ok := m[item]; !ok {
				existsInAllMaps = false
				break
			}
		}

		// If the current item exists in all of the maps, increment the count and print the item and its priority.
		if existsInAllMaps {
			position := unicode.ToLower(item) - rune('a') + 1
			if unicode.IsUpper(item) {
				position = position + 26
			}
			count += int(getPriority(item))

			fmt.Printf("The character '%c' only occurs once and has a priority of %n\n", item, int(position))
		}
	}
	return count
}

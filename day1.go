package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Iterate through the map of Elves to find the one carrying the most Calories.
	maxCalories, secondMaxCalories, thirdMaxCalories, elfCalories := 0, 0, 0, 0

	//var maxElf string

	// Define the path of the file to read.
	filePath := "day1input.txt"

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
	for scanner.Scan() {

		line := scanner.Text()

		if line != "" {
			//fmt.Println(len(line))
			calories, _ := strconv.Atoi(line)
			elfCalories += calories
			//fmt.Println("Adding", calories, "to elf", counter, "for a total of", elfCalories)
		} else {
			counter++
			if elfCalories > maxCalories {
				thirdMaxCalories = secondMaxCalories
				secondMaxCalories = maxCalories
				maxCalories = elfCalories
			} else if elfCalories > secondMaxCalories {
				thirdMaxCalories = secondMaxCalories
				secondMaxCalories = elfCalories
			} else if elfCalories > thirdMaxCalories {
				thirdMaxCalories = elfCalories
			}
			elfCalories = 0
			fmt.Println("Max Calories: ", maxCalories)
			fmt.Println("Second Max Calories: ", secondMaxCalories)
			fmt.Println("Third Max Calories: ", thirdMaxCalories)
		}
	}
	fmt.Println("we have", counter, "elves")

	// Handle last elf
	if elfCalories > maxCalories {
		if elfCalories > maxCalories {
			thirdMaxCalories = secondMaxCalories
			secondMaxCalories = maxCalories
			maxCalories = elfCalories
		} else if elfCalories > secondMaxCalories {
			thirdMaxCalories = secondMaxCalories
			secondMaxCalories = elfCalories
		} else if elfCalories > thirdMaxCalories {
			thirdMaxCalories = elfCalories
		}
	}

	fmt.Println("Max Calories: ", maxCalories)
	fmt.Println("Second Max Calories: ", secondMaxCalories)
	fmt.Println("Third Max Calories: ", thirdMaxCalories)

	fmt.Println("Total Calories: ", maxCalories+secondMaxCalories+thirdMaxCalories)

}

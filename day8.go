package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define the path of the file to read.
	filePath := "day8input.txt"
	const gridSize = 99

	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

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

	row := 0

	// Loop through the file line by line.
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			treeHeight, _ := strconv.Atoi(line[i : i+1])
			grid[row][i] = treeHeight
		}
		row++
	}

	//loop through grid and check visibility
	counter := 0
	maxTreeScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1 {
				// Shortcut for edges
				counter++
			} else {
				// Part 1
				if checkVisibility(grid, i, j) == 1 {
					counter++
				}
				// Part 2
				treeScore := checkTreeScore(grid, i, j)
				if treeScore > maxTreeScore {
					maxTreeScore = treeScore
				}
			}
		}
	}
	fmt.Println("trees visible: ", counter)
	fmt.Println("max tree score:", maxTreeScore)
}

func checkVisibility(grid [][]int, row int, col int) int {
	trees := 0

	// check up
	for i := row - 1; i >= 0; i-- {
		if grid[i][col] < grid[row][col] {
			trees++
		} else {
			// blocked, stop searching
			trees = 0
			break
		}
	}

	// found one, we are done
	if trees > 0 {
		return 1
	}

	// check down
	for i := row + 1; i < len(grid); i++ {
		if grid[i][col] < grid[row][col] {
			trees++
		} else {
			// blocked, stop searching
			trees = 0
			break
		}
	}

	// found one, we are done
	if trees > 0 {
		return 1
	}

	// check left
	for i := col - 1; i >= 0; i-- {
		if grid[row][i] < grid[row][col] {
			trees++
		} else {
			// blocked, stop searching
			trees = 0
			break
		}
	}

	// found one, we are done
	if trees > 0 {
		return 1
	}

	// check right
	for i := col + 1; i < len(grid[0]); i++ {
		if grid[row][i] < grid[row][col] {
			trees++
		} else {
			// blocked, stop searching
			trees = 0
			break
		}
	}

	// found one, we are done
	if trees > 0 {
		return 1
	}

	// found none
	return 0
}

func checkTreeScore(grid [][]int, row int, col int) int {
	var counterUp, counterDown, counterLeft, counterRight int

	// check up
	for i := row - 1; i >= 0; i-- {
		if grid[i][col] < grid[row][col] {
			counterUp++
		} else if grid[i][col] == grid[row][col] {
			counterUp++
			break
		} else {
			counterUp++
			break
		}
	}

	// check down
	for i := row + 1; i < len(grid); i++ {
		if grid[i][col] < grid[row][col] {
			counterDown++
		} else if grid[i][col] == grid[row][col] {
			counterDown++
			break
		} else {
			counterDown++
			break
		}
	}

	// check left
	for i := col - 1; i >= 0; i-- {
		if grid[row][i] < grid[row][col] {
			counterLeft++
		} else if grid[row][i] == grid[row][col] {
			counterLeft++
			break
		} else {
			counterLeft++
			break
		}
	}

	// check right
	for i := col + 1; i < len(grid[0]); i++ {
		if grid[row][i] < grid[row][col] {
			counterRight++
		} else if grid[row][i] == grid[row][col] {
			counterRight++
			break
		} else {
			counterRight++
			break
		}
	}

	return counterUp * counterDown * counterLeft * counterRight
}

//1763
// 671160

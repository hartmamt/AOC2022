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
	filePath := "day5input.txt"

	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		// If an error occurred, print it and return.
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Make nine stacks
	stacks := make([]Stack, 9)
	stacksPartTwo := make([]Stack, 9)

	// Create a new scanner for the file.
	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {

		line := scanner.Text()

		if counter < 8 {
			// Scan lines into crates
			for i := 0; i < 35; i++ {
				crate := line[i : i+3]
				column := i / 4
				// Push the crate onto the stack
				if crate != "   " {
					stacks[column].Load(crate)
					stacksPartTwo[column].Load(crate)
				}
				i = i + 3
			}
		} else if counter > 9 {
			// Now parse the moves
			parts := strings.Split(line, " ")
			if parts[0] == "move" {
				// How many creates to move
				num, _ := strconv.Atoi(parts[1])
				// Which stack to move from
				from, _ := strconv.Atoi(parts[3])
				// Which stack to move to
				to, _ := strconv.Atoi(parts[5])

				// This loop is for part one
				for i := 0; i < num; i++ {
					// move them crates one at a time (part one)
					if stacks[from-1].Len() > 0 {
						stacks[to-1].Push(stacks[from-1].Pop())
					}
				}

				// This loop is for part two
				// We need a temporary stack to hold the crates
				tempStack := Stack{}
				// Load them stacks into a stack
				for i := 0; i < num; i++ {
					if stacksPartTwo[from-1].Len() > 0 {
						tempStack.Push(stacksPartTwo[from-1].Pop())
					}
				}
				// Dump that stack onto the other stack
				for i := 0; i < num; i++ {
					stacksPartTwo[to-1].Push(tempStack.Pop())
				}

			}
		}
		counter++
	}

	// Print the answers

	fmt.Println()
	fmt.Println("Top of Stacks Part 1:")
	for i := 0; i < 9; i++ {
		fmt.Print(stacks[i].Peek())
	}

	fmt.Println()
	fmt.Println("Top of Stacks Part 2:")
	for i := 0; i < 9; i++ {
		fmt.Print(stacksPartTwo[i].Peek())
	}
}

// Stack is a simple stack data structure
type Stack struct {
	items []string
}

// Load adds an item to the bottom of the stack
func (s *Stack) Load(i string) {
	s.items = append([]string{i}, s.items...)
}

// Push adds an item to the top of the stack
func (s *Stack) Push(i string) {
	s.items = append(s.items, i)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top item from the stack without removing it
func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// Len returns the number of items in the stack
func (s *Stack) Len() int {
	return len(s.items)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
				}
				i = i + 3
			}
		} else if counter > 9 {
			// Split the string by spaces
			parts := strings.Split(line, " ")
			// Parse the move command
			if parts[0] == "move" {
				num, _ := strconv.Atoi(parts[1])
				from, _ := strconv.Atoi(parts[3])
				to, _ := strconv.Atoi(parts[5])

				fmt.Printf("Moving %d crates from stack %d to stack %d\n", num, from, to)

				for i := 0; i < num; i++ {
					// Pop the crates from the stack
					if stacks[from-1].Len() > 0 {
						stacks[to-1].Push(stacks[from-1].Pop())
					}
				}
			}
		}
		if counter == 8 {
			// Print the stacks
			fmt.Println()
			for i := 0; i < 9; i++ {
				fmt.Printf("Stack %d: ", i+1)
				for j := 0; j < stacks[i].Len(); j++ {
					fmt.Printf("%s ", stacks[i].items[j])
				}
				fmt.Println()
			}
			fmt.Println()
		}
		counter++
	}

	fmt.Println()
	fmt.Println("Ending Stacks:")
	for i := 0; i < 9; i++ {
		fmt.Println(stacks[i])
	}
	fmt.Println()
	fmt.Println("Top of Stacks:")
	for i := 0; i < 9; i++ {
		fmt.Print(stacks[i].Peek())
	}
	//[M][Q][T][P][G][L][L][D][N]
}

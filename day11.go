package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	number      int
	operation   string
	test        int
	throwTrue   int
	throwFalse  int
	items       []int
	inspections int
}

func main() {
	// Define the path of the file to read.
	filePath := "day11input.txt"

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

	// slice of monkeys
	monkeys := make([]Monkey, 0)

	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			// new monkey
			monkey := Monkey{}
			monkey.number = counter
			monkeys = append(monkeys, monkey)
		} else if strings.HasPrefix(line, "  Starting") {
			// starting items
			items := strings.Split(line[18:], ", ")
			// convert to ints
			for _, item := range items {
				i, _ := strconv.Atoi(item)
				monkeys[counter].items = append(monkeys[counter].items, i)
			}

		} else if strings.HasPrefix(line, "  Operation") {
			// operation
			monkeys[counter].operation = line[12:]
		} else if strings.HasPrefix(line, "  Test") {
			// test
			fmt.Println("TEST", line[21:])
			test, _ := strconv.Atoi(line[21:])
			monkeys[counter].test = test
		} else if strings.Contains(line, "throw") {
			// throw
			if strings.Contains(line, "true") {
				monkeys[counter].throwTrue, _ = strconv.Atoi(line[len(line)-1:])
			} else {
				monkeys[counter].throwFalse, _ = strconv.Atoi(line[len(line)-1:])
			}
		} else if line == "" {
			counter++
		}
	}
	//fmt.Println(len(monkeys))

	// 20 rounds
	for i := 0; i < 20; i++ {
		fmt.Println("Round", i+1)
		// loop through monkeys
		for c := 0; c < len(monkeys); c++ {
			//fmt.Println(monkey)
			//fmt.Println(c)
			for _, item := range monkeys[c].items {
				monkeys[c].inspections++
				// monkey still in play
				// get first item
				worry := monkeys[c].items[0]
				// parse operation
				operation := strings.Split(monkeys[c].operation, " ")
				//fmt.Println(operation)
				if operation[4] == "*" {
					//	fmt.Println("multiply")
					if operation[5] != "old" {
						value, _ := strconv.Atoi(operation[5])
						//		fmt.Println(item * value)
						worry = item * value
					} else {
						//		fmt.Println(item * item)
						worry = item * item
					}

				}
				if operation[4] == "+" {
					//	fmt.Println("add")
					if operation[5] != "old" {
						value, _ := strconv.Atoi(operation[5])
						//	fmt.Println(item + value)
						worry = item + value
					} else {
						//	fmt.Println(item + item)
						worry = item + item
					}
				}

				worry = worry / 3

				throwTo := 0
				if worry%monkeys[c].test == 0 {
					// throw true
					//	fmt.Println("throw true")
					throwTo = monkeys[c].throwTrue

				} else {
					// throw false
					//	fmt.Println("throw false")
					throwTo = monkeys[c].throwFalse
				}

				monkeys[throwTo].items = append(monkeys[throwTo].items, worry)
				monkeys[c].items = monkeys[c].items[1:]
			}

			fmt.Println("After round", i+1, c, "the monkeys are holding items with these worry levels:")
			for _, monkey := range monkeys {
				fmt.Println("Monkey:", monkey.number, monkey.items)
			}
		}
		// print inspections
		fmt.Println("After round", i+1, "the monkeys are holding items with these worry levels:")
		for _, monkey := range monkeys {
			fmt.Println("Monkey:", monkey.number, monkey.items)
		}
	}

	// print inspections
	// keep the two highest inspections
	inspectionsOne, inspectionsTwo := 0, 0
	for _, monkey := range monkeys {
		if monkey.inspections > inspectionsOne {
			inspectionsTwo = inspectionsOne
			inspectionsOne = monkey.inspections
		}

		fmt.Println("Monkey", monkey.number, "inspected", monkey.inspections, "items")
	}
	fmt.Println(inspectionsOne * inspectionsTwo)

	// 57348

}

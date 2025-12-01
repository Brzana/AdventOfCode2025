package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solver(input []string) {
	fmt.Println("Solver function called")
	
	numOfZeros := 0
	dial := 50

	for _, line := range input {

		line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }
        if len(line) < 2 {
            fmt.Println("Skipping malformed line:", line)
            continue
        }


		direction := line[:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting steps:", err)
			continue
		}

		if direction == "L" {
			dial -= steps
			outOfBounds(&dial, &numOfZeros)
		} else if direction == "R" {
			dial += steps
			outOfBounds(&dial, &numOfZeros)
		} else {
			fmt.Println("Invalid direction:", direction)
		}
	}
	fmt.Println("Number of times dial hit zero:", numOfZeros)
}

func outOfBounds(num *int, numOfZeros *int) {
	for *num < 0 {
		*numOfZeros++ 
		*num += 100
	}
	for *num > 99 {
		*numOfZeros++ 
		*num -= 100
	}
}
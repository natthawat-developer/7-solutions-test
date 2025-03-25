package main

import (
	"fmt"
	"strings"
)


func decodeSequence(encoded string) string {
	n := len(encoded) + 1
	numbers := make([]int, n)

	changed := true
	for changed {
		changed = false

		// Forward pass
		for i := range len(encoded) {
			if encoded[i] == 'R' && numbers[i+1] < numbers[i]+1 {
				numbers[i+1] = numbers[i] + 1
				changed = true
			} else if encoded[i] == '=' && numbers[i+1] != numbers[i] {
				numbers[i+1] = numbers[i]
				changed = true
			}
		}

		// Backward pass
		for i := len(encoded) - 1; i >= 0; i-- {
			if encoded[i] == 'L' && numbers[i] < numbers[i+1]+1 {
				numbers[i] = numbers[i+1] + 1
				changed = true
			} else if encoded[i] == '=' && numbers[i] != numbers[i+1] {
				numbers[i] = numbers[i+1]
				changed = true
			}
		}
	}

	// Normalize in a single pass
	minNum := numbers[0]
	for i := 1; i < n; i++ {
		if numbers[i] < minNum {
			minNum = numbers[i]
		}
	}
	for i := range numbers {
		numbers[i] -= minNum
	}

	// Convert to string
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
}

func main() {
	var encoded string
	fmt.Print("Enter encoded string: ")
	fmt.Scan(&encoded)
	fmt.Println("Decoded sequence:", decodeSequence(encoded))
}

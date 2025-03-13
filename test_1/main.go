package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func maxPathSum(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fileName := "hard.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var hardTriangle [][]int
	err = json.Unmarshal(data, &hardTriangle)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Max Path Sum from file:", maxPathSum(hardTriangle))
}

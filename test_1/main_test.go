package main

import "testing"

func TestMaxPathSum(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			[][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			237,
		},
		{
			[][]int{
				{7},
				{3, 8},
				{8, 1, 0},
				{2, 7, 4, 4},
				{4, 5, 2, 6, 5},
			},
			30,
		},
	}

	for _, tc := range testCases {
		result := maxPathSum(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}

package main

import "testing"

func TestDecodeSequence(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
	}

	for _, tc := range testCases {
		result := decodeSequence(tc.input)
		if result != tc.expected {
			t.Errorf("For input %s, expected %s but got %s", tc.input, tc.expected, result)
		}
	}
}
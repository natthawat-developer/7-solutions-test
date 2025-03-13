package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMeatWords(t *testing.T) {
	text := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
	expected := map[string]int{
		"t-bone":    4,
		"fatback":   1,
		"pastrami":  1,
		"pork":      1,
		"meatloaf":  1,
		"jowl":      1,
		"enim":      1,
		"bresaola":  1,
	}

	result := ExtractMeatWords(text)

	assert.Equal(t, expected, result, "Meat word count should be correct")
}

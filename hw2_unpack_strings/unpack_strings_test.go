package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSlash(t *testing.T) {
	assert.Equal(t, false, isSlash(rune('A')), "'A' => false")
	assert.Equal(t, false, isSlash(rune('z')), "'z' => false")
	assert.Equal(t, false, isSlash(rune('5')), "'5' => false")
	assert.Equal(t, true, isSlash(rune('\\')), "'\\' => true")
}

type dataStr struct {
	input  string
	output string
}

var testDataStr = []dataStr{
	{
		input:  "a4bc2d5e",
		output: "aaaabccddddde",
	},
	{
		input:  "abcd",
		output: "abcd",
	},
	{
		input:  "qwe\\4\\5",
		output: "qwe45",
	},
	{
		input:  "qwe\\45",
		output: "qwe44444",
	},
	{
		input:  "qwe\\\\5",
		output: "qwe\\\\\\\\\\",
	},
}

func TestUnpackString(t *testing.T) {
	for i := 0; i < len(testDataStr); i++ {
		expected := testDataStr[i].output
		actual, err := UnpackString(testDataStr[i].input)
		assert.Equal(t, expected, actual, "should be equal")
		assert.Nil(t, err)
	}

	val, err := UnpackString("45")
	assert.Equal(t, "", val, "should be equal")
	assert.NotNil(t, err)
}

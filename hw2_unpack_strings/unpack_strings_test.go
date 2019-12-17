package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber(rune('0')), "'0' => true")
	assert.Equal(t, true, isNumber(rune('9')), "'9' => true")
	assert.Equal(t, false, isNumber(rune('A')), "'A' => false")
	assert.Equal(t, false, isNumber(rune('\\')), "'A' => false")
}

func TestIsChar(t *testing.T) {
	assert.Equal(t, true, isChar(rune('A')), "'A' => true")
	assert.Equal(t, true, isChar(rune('z')), "'z' => true")
	assert.Equal(t, false, isChar(rune('5')), "'5' => false")
	assert.Equal(t, false, isChar(rune('\\')), "'\\' => false")
}

func TestIsSlash(t *testing.T) {
	assert.Equal(t, false, isSlash(rune('A')), "'A' => false")
	assert.Equal(t, false, isSlash(rune('z')), "'z' => false")
	assert.Equal(t, false, isSlash(rune('5')), "'5' => false")
	assert.Equal(t, true, isSlash(rune('\\')), "'\\' => true")
}

func TestMultiplyChar(t *testing.T) {
	assert.Equal(t, "ZZZ", multiplyChar(rune('Z'), 3), "should be equal")
	assert.Equal(t, "", multiplyChar(rune('Z'), 0), "should be equal")
	assert.Equal(t, "ZZZZZZZZZ", multiplyChar(rune('Z'), 9), "should be equal")
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

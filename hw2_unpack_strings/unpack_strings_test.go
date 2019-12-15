package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber([]rune("0")[0]), "'0' => true")
	assert.Equal(t, true, isNumber([]rune("9")[0]), "'9' => true")
	assert.Equal(t, false, isNumber([]rune("A")[0]), "'A' => false")
	assert.Equal(t, false, isNumber([]rune("\\")[0]), "'A' => false")
}

func TestIsChar(t *testing.T) {
	assert.Equal(t, true, isChar([]rune("A")[0]), "'A' => true")
	assert.Equal(t, true, isChar([]rune("z")[0]), "'z' => true")
	assert.Equal(t, false, isChar([]rune("5")[0]), "'5' => false")
	assert.Equal(t, false, isChar([]rune("\\")[0]), "'\\' => false")
}

func TestIsSlash(t *testing.T) {
	assert.Equal(t, false, isSlash([]rune("A")[0]), "'A' => false")
	assert.Equal(t, false, isSlash([]rune("z")[0]), "'z' => false")
	assert.Equal(t, false, isSlash([]rune("5")[0]), "'5' => false")
	assert.Equal(t, true, isSlash([]rune("\\")[0]), "'\\' => true")
}

func TestMultiplyChar(t *testing.T) {
	assert.Equal(t, "ZZZ", multiplyChar([]rune("Z")[0], 3), "should be equal")
	assert.Equal(t, "", multiplyChar([]rune("Z")[0], 0), "should be equal")
	assert.Equal(t, "ZZZZZZZZZ", multiplyChar([]rune("Z")[0], 9), "should be equal")
}

func TestUnpackString(t *testing.T) {
	val, err := UnpackString("a4bc2d5e")
	assert.Equal(t, "aaaabccddddde", val, "should be equal")
	assert.Nil(t, err)

	val, err = UnpackString("abcd")
	assert.Equal(t, "abcd", val, "should be equal")
	assert.Nil(t, err)

	val, err = UnpackString("45")
	assert.Equal(t, "", val, "should be equal")
	assert.NotNil(t, err)

	val, err = UnpackString("qwe\\4\\5")
	assert.Equal(t, "qwe45", val, "should be equal")
	assert.Nil(t, err)

	val, err = UnpackString("qwe\\45")
	assert.Equal(t, "qwe44444", val, "should be equal")
	assert.Nil(t, err)

	val, err = UnpackString("qwe\\\\5")
	assert.Equal(t, "qwe\\\\\\\\\\", val, "should be equal")
	assert.Nil(t, err)
}

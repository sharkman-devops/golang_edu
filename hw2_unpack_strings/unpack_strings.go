//Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
//
//* "a4bc2d5e" => "aaaabccddddde"
//* "abcd" => "abcd"
//* "45" => "" (некорректная строка)
//
//Дополнительное задание: поддержка escape - последовательности
//* `qwe\4\5` => `qwe45` (*)
//* `qwe\45` => `qwe44444` (*)
//* `qwe\\5` => `qwe\\\\\` (*)

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// UnpackString function for unpack string like: "a4bc2d5e" => "aaaabccddddde"
func UnpackString(inputStr string) (string, error) {
	result := ""

	for index, char := range inputStr {
		//fmt.Println(string(char))

		if index == 0 && !isSlash(char) && !isChar(char) {
			return "", errors.New("first symbol must be a char(a-z or A-Z) or '\\'")
		}

		var prevChar rune = 0
		if len(result) >= 1 {
			prevChar = []rune(result)[len(result)-1:][0]
		}

		if isChar(char) {
			if index >= 2 && isSlash(rune(inputStr[index-1])) && !isSlash(rune(inputStr[index-2])) {
				// \a => error
				return "", errors.New("wrong sequence")
			}

			if index == 1 && isSlash(rune(inputStr[index-1])) {
				// \a => error
				return "", errors.New("wrong sequence")
			}

			result += string(char)

		} else if isNumber(char) {
			if index >= 2 && isSlash(rune(inputStr[index-1])) && isSlash(rune(inputStr[index-2])) {
				// \\5 => \\\\\
				num, _ := strconv.Atoi(string(char))
				result += multiplyChar(prevChar, num-1)
			} else if isSlash(rune(inputStr[index-1])) {
				// \5 => 5
				result += string(char)
			} else {
				// a5 => aaaaa
				num, _ := strconv.Atoi(string(char))
				result += multiplyChar(prevChar, num-1)
			}
		} else if isSlash(char) {
			if index >= 1 && isSlash(rune(inputStr[index-1])) {
				// \\ => \
				result += string(char)
			}

		}

	}
	return result, nil
}

func isNumber(char rune) bool {
	if char >= rune('0') && char <= rune('9') {
		return true
	}
	return false
}

func isChar(char rune) bool {
	if char >= rune('A') &&
		char <= rune('Z') ||
		char >= rune('a') &&
			char <= rune('z') {
		return true
	}
	return false
}

func isSlash(char rune) bool {
	if char == rune('\\') {
		return true
	}
	return false
}

func multiplyChar(char rune, multiplier int) (result string) {
	//fmt.Println(char, multiplier)
	for i := 0; i < multiplier; i++ {
		result += string(char)
	}
	return result
}

func main() {
	val, err := UnpackString("z4\\\\bc2d5e\\\\5")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(val)
	}
}

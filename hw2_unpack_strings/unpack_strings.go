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
	"strings"
	"unicode"
)

// UnpackString function for unpack string like: "a4bc2d5e" => "aaaabccddddde"
func UnpackString(inputStr string) (string, error) {
	result := ""

	for index, char := range inputStr {
		//fmt.Println(string(char))

		if index == 0 && !isSlash(char) && !unicode.IsLetter(char) {
			return "", errors.New("first symbol must be a char(a-z or A-Z) or '\\'")
		}

		prevChar := ""
		if len(result) >= 1 {
			prevChar = result[len(result)-1:]
		}

		switch {
		case unicode.IsLetter(char):
			if index >= 2 && isSlash(rune(inputStr[index-1])) && !isSlash(rune(inputStr[index-2])) {
				// \a => error
				return "", errors.New("wrong sequence")
			}

			if index == 1 && isSlash(rune(inputStr[index-1])) {
				// \a => error
				return "", errors.New("wrong sequence")
			}

			result += string(char)
		case unicode.IsDigit(char):
			if index >= 2 && isSlash(rune(inputStr[index-1])) && isSlash(rune(inputStr[index-2])) {
				// \\5 => \\\\\
				num, _ := strconv.Atoi(string(char))
				result += strings.Repeat(prevChar, num-1)
			} else if isSlash(rune(inputStr[index-1])) {
				// \5 => 5
				result += string(char)
			} else {
				// a5 => aaaaa
				num, _ := strconv.Atoi(string(char))
				result += strings.Repeat(prevChar, num-1)
			}
		case isSlash(char):
			if index >= 1 && isSlash(rune(inputStr[index-1])) {
				// \\ => \
				result += string(char)
			}
		}
	}
	return result, nil
}

func isSlash(char rune) bool {
	if char == rune('\\') {
		return true
	}
	return false
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

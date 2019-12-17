//Частотный анализ
//Цель: Напиcать функцию, принимающую на вход строку с текстом и возвращающую слайс с 10 самыми частовстречающимеся в тексте словами.
//Если есть более 10 самых частотых слов (например 15 разных слов встречаются ровно 133 раза, остальные < 100), можно вернуть любые 10 из самых частотных.
//Словоформы не учитываем. "нога", "ногу", "ноги" - это разные слова. Слово с большой и маленькой буквы можно считать за разные слова.
//"Нога" и "нога" - это разные слова. Знаки препиания можно считать "буквами" слова или отдельными словами.
//"-" (тире) - это отдельное слово. "нога," и "нога" - это разные слова. Пример: "cat and dog one dog two cats and one man".
//"dog", "one", "and" - встречаются два раза, это топ-3. Задание со звездочкой (*): учитывать большие/маленьгие буквы и знаки препинания.
//"Нога" и "нога" - это одинаковые слова, "нога," и "нога" - это одинаковые слова, "-" (тире) - это не слово.

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Top10 return top10 words of string
func Top10(input string) []string {
	lowerInput := strings.ToLower(input)
	cleanInput := ""
	for _, char := range lowerInput {
		if char >= []rune("a")[0] && char <= []rune("z")[0] || char == []rune(" ")[0] {
			cleanInput += string(char)
		}
	}

	spaceSplitted := strings.Split(cleanInput, " ")
	//fmt.Println(spaceSplitted)
	wordsFreq := map[string]int{}
	for _, word := range spaceSplitted {
		wordsFreq[word]++
	}

	freq := make([]int, len(wordsFreq))
	for _, val := range wordsFreq {
		freq = append(freq, val)
	}

	sort.Ints(freq)
	var top10Freq []int
	if len(freq) >= 10 {
		top10Freq = freq[len(freq)-10:]
	} else {
		top10Freq = freq[:]
	}

	top10 := []string{}
	for _, value := range top10Freq {
		for key, val := range wordsFreq {
			if val == value {
				top10 = append(top10, key)
				delete(wordsFreq, key)
			}
		}
	}

	//fmt.Println(top10Freq)
	//fmt.Println(wordsFreq)
	return top10
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris gravida mauris a lectus cursus, venenatis fringilla purus aliquet. Curabitur et elementum nulla. Nulla facilisi. Etiam facilisis viverra consectetur. Aliquam consequat turpis id nisi dapibus, sed commodo lacus luctus. Morbi ac lacinia massa. Cras enim metus, posuere at aliquet volutpat, auctor sit amet erat. Sed vel ex aliquet justo laoreet aliquet. Donec ultricies leo tellus. Donec vitae bibendum lectus. Lorem."
	fmt.Println(Top10(text))
}

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
	"unicode"
)

// Top10 return top10 words of string
func Top10(input string) []string {
	lowerInput := strings.ToLower(input)
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	cleanInputSlice := strings.FieldsFunc(lowerInput, f)

	type wordFreq struct {
		word string
		freq int
	}
	wordsFreqSlice := []wordFreq{}

	type indexAndCount struct {
		index int
		count int
	}
	wordsFreqMap := make(map[string]indexAndCount)

	for _, word := range cleanInputSlice {
		_, ok := wordsFreqMap[word]
		if ok {
			wordsFreqMap[word] = indexAndCount{index: wordsFreqMap[word].index, count: wordsFreqMap[word].count + 1}
			wordsFreqSlice[wordsFreqMap[word].index] = wordFreq{word: word, freq: wordsFreqMap[word].count}
		} else {
			wordsFreqMap[word] = indexAndCount{index: len(wordsFreqSlice), count: 1}
			wordsFreqSlice = append(wordsFreqSlice, wordFreq{word, 1})
		}
	}

	sort.Slice(wordsFreqSlice, func(i, j int) bool { return wordsFreqSlice[i].freq > wordsFreqSlice[j].freq })

	top10 := []string{}
	maxLen := 10
	if len(wordsFreqSlice) < 10 {
		maxLen = len(wordsFreqSlice)
	}

	for i := 0; i < maxLen; i++ {
		top10 = append(top10, wordsFreqSlice[i].word)
	}

	return top10
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris gravida mauris a lectus cursus, venenatis fringilla purus aliquet. Curabitur et elementum nulla. Nulla facilisi. Etiam facilisis viverra consectetur. Aliquam consequat turpis id nisi dapibus, sed commodo lacus luctus. Morbi ac lacinia massa. Cras enim metus, posuere at aliquet volutpat, auctor sit amet erat. Sed vel ex aliquet justo laoreet aliquet. Donec ultricies leo tellus. Donec vitae bibendum lectus. Lorem."
	fmt.Println(Top10(text))
}

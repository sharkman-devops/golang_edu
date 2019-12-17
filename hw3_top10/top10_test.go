package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dataStr struct {
	input  string
	output []string
}

var testDataStr = []dataStr{
	{
		input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris gravida mauris a lectus cursus, venenatis fringilla purus aliquet. Curabitur et elementum nulla. Nulla facilisi. Etiam facilisis viverra consectetur. Aliquam consequat turpis id nisi dapibus, sed commodo lacus luctus. Morbi ac lacinia massa. Cras enim metus, posuere at aliquet volutpat, auctor sit amet erat. Sed vel ex aliquet justo laoreet aliquet. Donec ultricies leo tellus. Donec vitae bibendum lectus. Lorem.",
		output: []string{"aliquet", "lorem", "lectus", "nulla", "consectetur", "donec", "mauris", "sed", "sit", "amet"},
	},
	{
		input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		output: []string{"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit"},
	},
	{
		input:  "",
		output: []string{""},
	},
}

func TestTop10(t *testing.T) {
	for i := 0; i < len(testDataStr); i++ {
		expected := testDataStr[i].output
		sort.Strings(expected)
		actual := Top10(testDataStr[i].input)
		sort.Strings(actual)

		assert.Equal(t, expected, actual, "should be equal")
	}
}

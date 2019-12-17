package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop10(t *testing.T) {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris gravida mauris a lectus cursus, venenatis fringilla purus aliquet. Curabitur et elementum nulla. Nulla facilisi. Etiam facilisis viverra consectetur. Aliquam consequat turpis id nisi dapibus, sed commodo lacus luctus. Morbi ac lacinia massa. Cras enim metus, posuere at aliquet volutpat, auctor sit amet erat. Sed vel ex aliquet justo laoreet aliquet. Donec ultricies leo tellus. Donec vitae bibendum lectus. Lorem."
	expected := []string{"aliquet", "lorem", "lectus", "nulla", "consectetur", "donec", "mauris", "sed", "sit", "amet"}
	sort.Strings(expected)
	actual := Top10(text)
	sort.Strings(actual)
	assert.Equal(t, expected, actual, "should be equal")

	text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	expected = []string{"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit"}
	sort.Strings(expected)
	actual = Top10(text)
	sort.Strings(actual)
	assert.Equal(t, expected, actual, "should be equal")

	text = ""
	expected = []string{""}
	actual = Top10(text)
	assert.Equal(t, expected, actual, "should be equal")

}

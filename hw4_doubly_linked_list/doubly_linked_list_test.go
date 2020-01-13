package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList_Len(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	assert.Equal(t, 0, doublyLinkedList.Len())

	doublyLinkedList.PushBack(1)
	assert.Equal(t, 1, doublyLinkedList.Len())

	doublyLinkedList.Remove(*doublyLinkedList.Last())
	assert.Equal(t, 0, doublyLinkedList.Len())
}

func TestDoublyLinkedList_PushBack(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.PushBack(1)
	firstItem := doublyLinkedList.First()
	lastItem := doublyLinkedList.Last()
	assert.Equal(t, doublyLinkedList, DoublyLinkedList{len: 1, first: firstItem, last: lastItem})
	itm := *doublyLinkedList.first
	assert.Equal(t, Item{value: 1, prev: nil, next: nil}, itm)
	itm = *doublyLinkedList.last
	assert.Equal(t, Item{value: 1, prev: nil, next: nil}, itm)

	doublyLinkedList.PushBack(2)
	firstItem = doublyLinkedList.First()
	lastItem = doublyLinkedList.Last()
	assert.Equal(t, doublyLinkedList, DoublyLinkedList{len: 2, first: firstItem, last: lastItem})
	itm = *doublyLinkedList.first
	assert.Equal(t, Item{value: 1, prev: nil, next: doublyLinkedList.last}, itm)
	itm = *doublyLinkedList.last
	assert.Equal(t, Item{value: 2, prev: doublyLinkedList.first, next: nil}, itm)
}

func TestDoublyLinkedList_PushFront(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.PushFront("BAR")
	firstItem := doublyLinkedList.First()
	lastItem := doublyLinkedList.Last()
	assert.Equal(t, doublyLinkedList, DoublyLinkedList{len: 1, first: firstItem, last: lastItem})
	itm := *doublyLinkedList.first
	assert.Equal(t, Item{value: "BAR", prev: nil, next: nil}, itm)
	itm = *doublyLinkedList.last
	assert.Equal(t, Item{value: "BAR", prev: nil, next: nil}, itm)

	doublyLinkedList.PushFront("FOO")
	firstItem = doublyLinkedList.First()
	lastItem = doublyLinkedList.Last()
	assert.Equal(t, doublyLinkedList, DoublyLinkedList{len: 2, first: firstItem, last: lastItem})
	itm = *doublyLinkedList.first
	assert.Equal(t, Item{value: "FOO", prev: nil, next: doublyLinkedList.Last()}, itm)
	itm = *doublyLinkedList.last
	assert.Equal(t, Item{value: "BAR", prev: doublyLinkedList.First(), next: nil}, itm)

}

func TestDoublyLinkedList_Remove(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	itm := Item{prev: nil, next: nil, value: "foo"}
	err := doublyLinkedList.Remove(itm)
	assert.EqualError(t, err, "no items to delete")

	doublyLinkedList.PushBack("foo")
	err = doublyLinkedList.Remove(itm)
	assert.NoError(t, err)
	assert.Equal(t, 0, doublyLinkedList.Len())

	doublyLinkedList.PushBack("bar")
	err = doublyLinkedList.Remove(itm)
	assert.EqualError(t, err, "item not found")
	assert.Equal(t, 1, doublyLinkedList.Len())

	doublyLinkedList.PushBack("bar1")
	doublyLinkedList.PushBack("bar2")
	doublyLinkedList.PushBack("foo")
	doublyLinkedList.PushBack("bar3")
	err = doublyLinkedList.Remove(itm)
	assert.EqualError(t, err, "item not found")
	assert.Equal(t, 5, doublyLinkedList.Len())

	err = doublyLinkedList.Remove(*doublyLinkedList.Last().Prev())
	assert.NoError(t, err)
	assert.Equal(t, 4, doublyLinkedList.Len())
	assert.Equal(t, "bar", doublyLinkedList.First().Value())
	assert.Equal(t, "bar1", doublyLinkedList.First().Next().Value())
	assert.Equal(t, "bar2", doublyLinkedList.First().Next().Next().Value())
	assert.Equal(t, "bar3", doublyLinkedList.First().Next().Next().Next().Value())
}

func TestDoublyLinkedList_First(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	assert.Nil(t, doublyLinkedList.First())

	doublyLinkedList.PushBack("FOO")
	assert.Equal(t, &Item{value: "FOO", prev: nil, next: nil}, doublyLinkedList.First())
	doublyLinkedList.PushFront("BAR")
	assert.Equal(t, &Item{value: "BAR", prev: nil, next: doublyLinkedList.Last()}, doublyLinkedList.First())
}

func TestDoublyLinkedList_Last(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	assert.Nil(t, doublyLinkedList.Last())

	doublyLinkedList.PushBack("FOO")
	assert.Equal(t, &Item{value: "FOO", prev: nil, next: nil}, doublyLinkedList.Last())
	doublyLinkedList.PushFront("BAR")
	assert.Equal(t, &Item{value: "FOO", prev: doublyLinkedList.First(), next: nil}, doublyLinkedList.Last())
}

func TestItem_Next(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.PushBack("WORD")
	assert.Nil(t, doublyLinkedList.First().Next())

	doublyLinkedList.PushFront("HELLO")
	itm := doublyLinkedList.First()
	assert.Equal(t, &Item{next: nil, prev: doublyLinkedList.First(), value: "WORD"}, itm.Next())
	assert.Nil(t, itm.Next().Next())
}

func TestItem_Prev(t *testing.T) {
	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.PushBack("WORD")
	assert.Nil(t, doublyLinkedList.First().Prev())

	doublyLinkedList.PushFront("HELLO")
	itm := doublyLinkedList.Last()
	assert.Equal(t, &Item{next: itm, prev: nil, value: "HELLO"}, itm.Prev())
	assert.Nil(t, itm.Prev().Prev())
}

func TestItem_Value(t *testing.T) {
	emptyItem := Item{}
	assert.Nil(t, emptyItem.Value())

	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.PushBack("WORD")
	doublyLinkedList.PushFront("HELLO")
	itm := doublyLinkedList.First()
	assert.Equal(t, "HELLO", itm.Value())
	assert.Equal(t, "WORD", itm.Next().Value())
}

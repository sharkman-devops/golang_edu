package main

import (
	"errors"
	"fmt"
)

// DoublyLinkedList - main struct
type DoublyLinkedList struct {
	len   int
	first *Item
	last  *Item
}

// Len of a Doubly Linked List
func (d DoublyLinkedList) Len() int {
	return d.len
}

// First - returns the first item or nil
func (d DoublyLinkedList) First() *Item {
	return d.first
}

// Last - returns the last item or nil
func (d DoublyLinkedList) Last() *Item {
	return d.last
}

// PushFront - inserts item before the first item
func (d *DoublyLinkedList) PushFront(v interface{}) {
	d.len++
	firstItem := d.First()
	newItem := Item{
		value: v,
		prev:  nil,
		next:  firstItem,
	}
	if firstItem == nil {
		d.last = &newItem
	}
	d.first = &newItem

	if firstItem != nil {
		firstItem.prev = &newItem
	}
}

// PushBack - inserts item after the last item
func (d *DoublyLinkedList) PushBack(v interface{}) {
	d.len++
	lastItem := d.Last()
	newItem := Item{
		value: v,
		prev:  lastItem,
		next:  nil,
	}
	if lastItem == nil {
		d.first = &newItem
	}

	d.last = &newItem

	if lastItem != nil {
		newItem.prev = lastItem
		lastItem.next = &newItem
	}
}

// Remove item from a Doubly Linked List, returns nil or error
func (d *DoublyLinkedList) Remove(i Item) error {
	switch d.Len() {
	case 0:
		return errors.New("no items to delete")
	case 1:
		if i.prev != nil || i.next != nil {
			return errors.New("item not found")
		}
		if i.value != d.First().Value() {
			return errors.New("item not found")
		}

		d.first = nil
		d.last = nil
		d.len--
		return nil
	}

	switch i {
	case *d.First():
		d.first = i.next
		i.next.prev = nil
	case *d.Last():
		d.last = i.prev
		i.prev.next = nil
	default:
		if i.prev == nil || i.next == nil {
			return errors.New("item not found")
		}
		if *i.prev.next != i {
			return errors.New("item not found")
		}
		if *i.next.prev != i {
			return errors.New("item not found")
		}

		i.prev.next = i.next
		i.next.prev = i.prev
	}

	d.len--
	return nil
}

// Item - item of a Doubly Linked List
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

// Value - returns value of item
func (i Item) Value() interface{} {
	return i.value
}

// Next - returns the next item or nil
func (i Item) Next() *Item {
	return i.next
}

// Prev - returns the previous item or nil
func (i Item) Prev() *Item {
	return i.prev
}

func main() {
	dll := DoublyLinkedList{}
	dll.PushFront(2)
	dll.PushFront(1)
	dll.PushBack(3)
	dll.PushBack(4)
	dll.PushBack(5)
	dll.PushFront(0)

	fmt.Print("Expected:012345 Got:")
	itm := dll.First()
	for ; itm != nil; itm = itm.Next() {
		fmt.Print(itm.Value())
	}
	fmt.Println()

	dll.Remove(*dll.First())
	dll.Remove(*dll.Last().Prev().Prev())
	dll.Remove(*dll.Last())

	fmt.Print("Expected:124 Got:")
	itm = dll.First()
	for ; itm != nil; itm = itm.Next() {
		fmt.Print(itm.Value())
	}
	fmt.Println()
}

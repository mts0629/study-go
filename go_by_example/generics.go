package main

import "fmt"

// Generic function to get a slice index
// For slices of any `comparable` type
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	// Not found
	return -1
}

// Singly-linked list for generic types
type List[T any] struct {
	head, tail *element[T]
}

// Element of the list
type element[T any] struct {
	next *element[T]
	val  T
}

// Push method of the list
func (lst *List[T]) Push(v T) {
	// Need to keep type parameters as like `List[T]`
	if lst.tail == nil {
		// First element
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements method of the list,
// return all the elements as a slice
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "foobar"}

	// Get an element's index from a slice of string
	fmt.Println("index of foobar:", SlicesIndex(s, "foobar"))
	// Specify the type explicitly
	fmt.Println("index of foobar:", SlicesIndex[[]string, string](s, "foobar"))

	// Get an element's index from a slice of int
	var i = []int{1, 2, 3}
	fmt.Println("index of 2:", SlicesIndex(i, 2))

	// Operate an integer list
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}

// Generics
package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, e E) int {
	for i := range s {
		if s[i] == e {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *Element[T]
}

type Element[T any] struct {
	next *Element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &Element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &Element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems = []T{}
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"a", "b", "c"}
	fmt.Printf("Index: %v\n", SlicesIndex(s, "b"))

	lst := List[int]{}

	lst.Push(1)
	lst.Push(3)
	lst.Push(5)
	lst.Push(7)

	fmt.Println("List: ", lst.AllElements())
}

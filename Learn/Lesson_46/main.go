// sorting by func

package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func main() {
	fruits := []string{"aaa", "a", "aa"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	people := []Person{
		Person{name: "c", age: 15},
		Person{name: "a", age: 5},
		Person{name: "b", age: 10},
	}

	ageCpm := func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	}

	slices.SortFunc(people, ageCpm)
	fmt.Println(people)
}

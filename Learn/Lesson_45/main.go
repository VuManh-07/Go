// sorting
package main

import (
	"fmt"
	"slices"
)

func main() {
	str := []string{"d", "a", "a", "c", "b"}
	slices.Sort(str)
	fmt.Println(str)

	ints := []int{2, 4, 6, 10}
	slices.Sort(ints)
	fmt.Println(ints)

	iSort := slices.IsSorted(ints)
	fmt.Println(iSort)
}

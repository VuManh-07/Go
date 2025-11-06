// Recover
package main

import "fmt"

func main() {
	safeDivide := func(a, b int) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Khac phuc", r)
			}
		}()

		fmt.Println(a / b)
	}

	safeDivide(10, 0)
}

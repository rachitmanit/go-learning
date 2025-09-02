package main

import (
	"fmt"
)

func do(a [3]int, b []int) []int {
	// a = b   // SYNTAX ERROR

	a[0] = 4 // modifies local copy of array (does not affect w)
	b[0] = 3 // modifies slice (affects x outside)

	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42
	copy(c, b) // Copies min(len(c), len(b)) elements from b to c. Here it copies only 3 elements from b into c.

	return c
}

func main() {
	// Definition difference between array and slice.
	w := [...]int{1, 2, 3} // array of len(3)
	x := []int{0, 0, 0}    // slice of len(3)

	y := do(w, x)

	fmt.Println(w, x, y) // [1 2 3] [3 0 0] [3 0 0 42]

	// Definition 2 of array
	arr := [4]int{1, 2, 3} // array of len(4) but initialised with 3 values. 4th value is zero as default.
	fmt.Println(arr)       // [1 2 3 0]
}

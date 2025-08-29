package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var val float64
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Println("No elements were provided.")
		return
	}

	fmt.Printf("Average: %.2f, Total Numbers: %d", sum/float64(n), n)
}

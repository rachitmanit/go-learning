package main

import "fmt"

func main() {
	str1 := "elite"
	str2 := "èlite"

	fmt.Printf("Type: %8T Value: %[1]v\n", str1)
	fmt.Printf("Type: %8T Value: %[1]v\n", str2)

	fmt.Println("-----------------------")

	// 1. rune (int32 in length) type conversion
	fmt.Printf("%8T %[1]v\n", []rune(str1))
	fmt.Printf("%8T %[1]v\n", []rune(str2))

	fmt.Println("-----------------------")

	// 2. byte (int8 in length) type conversion. See how byte handles unicode characters
	// It uses an additional space to accomodate it.
	fmt.Printf("%8T %[1]v\n", []byte(str1))
	fmt.Printf("%8T %[1]v\n", []byte(str2))

	fmt.Println("-----------------------")

	// 3. NOTE: Length difference with accent character in normal strings length function.
	fmt.Printf("Len str1: %d\n", len(str1))
	fmt.Printf("Len str2: %d\n", len(str2))

}

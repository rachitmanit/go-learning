package main

import "fmt"

func example1() {
	var m map[string]int
	p := make(map[string]int, 0)

	println(m == nil) // TRUE
	println(p == nil) // FALSE

	// println(m == p)   // SYNTAX ERROR. Map can only be compared to "nil"

	a := m["and"]
	b := p["and"]
	println(a, b) // Prints 0, 0. Can search in empty and nil maps. It will return DEFAULT values.

	p["and"] += 1
	// m["and"] += 1 // PANIC. Because m is still nil and is not initiliased to a valid memory. Error: "panic: assignment to entry in nil map"

	m = p
	m["and"]++
	println(m["and"], p["and"]) // Prints 2,2. As m's descriptor is now referencing to p. And they both operate on shared memory.
}

func example2() {
	p := map[string]int{}
	println(p == nil) // FALSE. Because it is initialied using shorthand initialisation of zero size map

	a := p["and"]
	println(a) // Prints 0

	b, ok := p["and"]
	println(b, ok) // Prints 0, false. This false is used to check if value exists or not in map.

	p["and"]++ // Adds a key "and" and increments it by 1

	if _, ok := p["and"]; ok {
		println("Exists!")
	}
}

func main() {
	example1()

	fmt.Println("------------------------------------------------------")

	example2()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := map[string]int{}

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Printf("Unique Words: %v", len(words))

	type kv struct {
		k string
		v int
	}

	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].v > ss[j].v
	})

	for _, word := range ss {
		fmt.Println(word.k, "appears", word.v, "times")
	}
}

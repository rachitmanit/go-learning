package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Run command: cat .\input.txt | go run . John Henry

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Arguments are not enough. Should be two\n")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := map[string]int{}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		counts[s.Text()]++
	}

	for word, count := range counts {
		fmt.Printf("%q: %d\n", word, count)
	}
}

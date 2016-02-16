package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if line := s.Text(); re.MatchString(line) {
			fmt.Printf("%s\n", line)
		}
	}
}

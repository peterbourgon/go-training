package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		stderr = flag.Bool("stderr", false, "output to stderr instead of stdout")
		sz     = flag.Bool("sz", false, "print size at end")
	)
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s [options] <filename>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	dst := os.Stdout
	if *stderr {
		dst = os.Stderr
	}

	n, err := io.Copy(dst, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}

	if *sz {
		fmt.Fprintf(dst, "%d\n", n)
	}
}

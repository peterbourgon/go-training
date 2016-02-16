package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			abs, err := filepath.Abs(path)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", abs)
		}
		return nil
	})
}

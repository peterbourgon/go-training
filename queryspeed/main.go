package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	c := make(chan result, len(os.Args[1:]))
	for _, arg := range os.Args[1:] {
		go func(url string) { c <- get(url) }(arg)
	}
	for i := 0; i < cap(c); i++ {
		r := <-c
		fmt.Printf("%s: took %s, code %d, err %v\n", r.url, r.took, r.code, r.err)
	}
}

type result struct {
	url  string
	took time.Duration
	code int
	err  error
}

func get(url string) result {
	begin := time.Now()
	resp, err := http.Get(url)
	code := 0
	if err == nil {
		code = resp.StatusCode
		resp.Body.Close()
	}
	return result{
		url:  url,
		took: time.Since(begin),
		code: code,
		err:  err,
	}
}

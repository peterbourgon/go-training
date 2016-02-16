package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		target = flag.String("target", "google.com", "host to proxy to")
	)
	flag.Parse()

	http.HandleFunc("/", h(*target))
	log.Printf("listening on %s", *listen)
	http.ListenAndServe(*listen, nil)
}

func h(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.URL.Host = target
		if req.URL.Scheme == "" {
			req.URL.Scheme = "http"
		}
		log.Printf("%s => %s", r.URL, req.URL)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
}

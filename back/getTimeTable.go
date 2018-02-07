package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
)

type go struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
	To   []int `json:"to"`
}

type return struct {
	Fast []int `json:"fast"`
	From []int `json:"from"`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

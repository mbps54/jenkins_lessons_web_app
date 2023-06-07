package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from lesson 10!")
  })
  err := http.ListenAndServe("0.0.0.0:8081", nil)
  if err != nil {
    panic(err)
  }
}

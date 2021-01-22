package main

import (
  "fmt"
  "os"
  "log"
  "net/http"
)

func handler( w http.ResponseWriter, r *http.Request) {
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }

  fmt.Fprintln(w, "hostname:", name)
}

func main() {
  fmt.Fprintln(os.Stdout, "start Golange server...")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

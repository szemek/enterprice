package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", hello)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "3000"
  }
  hostname, _ := os.Hostname()
  fmt.Printf("listening on %s:%s...\n", hostname, port)
  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
    panic(err)
  }
}

func hello(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "hello, world")
}

package main

import (
  "fmt"
  "net/http"
  "os"
  "regexp"
)

func main() {
  http.HandleFunc("/", router)

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

func router(res http.ResponseWriter, req *http.Request) {
  if match("/hello", req.URL.Path) {
    fmt.Fprintln(res, "hello, world")
  } else {
    fmt.Fprintln(res, ".")
  }
}

func match(pattern string, s string) (matched bool) {
  matched, _ = regexp.MatchString(pattern, s)
  return matched
}

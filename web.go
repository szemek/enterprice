package main

import (
  "fmt"
  "net/http"
  "os"
  "regexp"
  "io/ioutil"
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
  if match("^/$", req.URL.Path) {
    body, _ := ioutil.ReadFile("./public/index.html")
    fmt.Fprintln(res, string(body))
  } else if match("^/[js|css]/*", req.URL.Path) {
    http.ServeFile(res, req, "./public/" + req.URL.Path)
  } else {
    fmt.Fprintln(res, "")
  }
}

func match(pattern string, s string) (matched bool) {
  matched, _ = regexp.MatchString(pattern, s)
  return matched
}

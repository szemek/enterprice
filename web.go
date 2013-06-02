package main

import (
  "fmt"
  "net/http"
  "os"
  "regexp"
  "io/ioutil"
  "time"
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

type Route struct {
  Method, Path string
}

func api(res http.ResponseWriter, req *http.Request) {
  method := req.Method
  path := req.URL.Path
  route := Route{method, path}

  switch route {
  case Route{"GET", "/products"}:
    fmt.Fprintln(res, "products")
  }

  fmt.Fprintln(res, req.Method)
}

func log(req *http.Request) {
  now := time.Now().Format(time.RFC3339)
  method := req.Method
  path := req.URL.Path
  address := req.RemoteAddr
  fmt.Printf("%s %s \"%s\" from %s\n", now, method, path, address)
}

func router(res http.ResponseWriter, req *http.Request) {
  log(req)

  if match("^/$", req.URL.Path) {
    body, _ := ioutil.ReadFile("./public/index.html")
    fmt.Fprintln(res, string(body))
  } else if match("^/[js|css]/*", req.URL.Path) {
    http.ServeFile(res, req, "./public/" + req.URL.Path)
  } else {
    api(res, req)
  }
}

func match(pattern string, s string) (matched bool) {
  matched, _ = regexp.MatchString(pattern, s)
  return matched
}

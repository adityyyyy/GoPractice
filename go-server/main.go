package main

import(
  "fmt"
  "log"
  "net/http"
)

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandlerFunc("/form", formHandler)
  http.HandlerFunc("/hello", helloHandler)

  fmt.Printf("Starting server at port: 8080\n")
  if
}

package main

import (
  "net/http"
  "html/template"
  )

type Page struct {
  Title string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/main.html")
  p := &Page{Title: "test"}
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
    })
  panic(http.ListenAndServe(":8080", nil))
}

package main

import (
  "net/http"
  "html/template"
  )

type Page struct {
  Title string
}

type Staff struct {
  FName string
  LName string
  Facility string
  EMail string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/main.html")
  s := &Staff{FName: "Brian", LName: "Lawrence", Facility: "Hospital 1", EMail: "b@b.net"}
  t.Execute(w, s)
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
    })
  panic(http.ListenAndServe(":8080", nil))
}

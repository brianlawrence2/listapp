package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

type Staff struct {
	StaffMembers []StaffMember
}

type StaffMember struct {
	FName    string
	LName    string
	Facility string
	Position string
	EMail    string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/main.html")
	sm1 := StaffMember{FName: "Brian", LName: "Lawrence", Facility: "Hospital 1", Position: "CEO", EMail: "b@b.net"}
	sm2 := StaffMember{FName: "Clark", LName: "Lawrence", Facility: "Hospital 2", Position: "CFO", EMail: "c@c.net"}
	s := &Staff{StaffMembers: []StaffMember{sm1, sm2}}
	t.Execute(w, s)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/edit.html")
	sm1 := StaffMember{FName: "Brian", LName: "Lawrence", Facility: "Hospital 1", Position: "CEO", EMail: "b@b.net"}
	sm2 := StaffMember{FName: "Clark", LName: "Lawrence", Facility: "Hospital 2", Position: "CFO", EMail: "c@c.net"}
	s := &Staff{StaffMembers: []StaffMember{sm1, sm2}}
	t.Execute(w, s)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	panic(http.ListenAndServe(":8080", nil))
}

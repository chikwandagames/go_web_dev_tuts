package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	// func ParseFiles(filenames ...string) (*Template, error)
	// tpl, _ = template.ParseFiles("index1.html")
	// tpl, _ = template.ParseFiles("data1/index2.html")
	// tpl, _ = template.ParseFiles("data1/data2/index3.html")
	// tpl, _ = template.ParseFiles("../index4.html")

	// Alternatively, we could use the ParseFiles directly
	tpl, _ = tpl.ParseFiles("../index4.html")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

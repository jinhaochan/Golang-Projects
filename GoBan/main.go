package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p := loadAllData()
	fmt.Println(p)
	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {

	http.HandleFunc("/", makeHandler(viewHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)

}

package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

func main() {
	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()
	// start the web server
	http.ListenAndServe(":8097", nil)
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(
		func() {
			t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		})
	t.templ.Execute(w, r)
}

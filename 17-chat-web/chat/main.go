package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

func main() {
	gomniauth.SetSecurityKey("831989856590-onu4to9s0aucvqt9pe3socp172ktfhu4.apps.googleusercontent.com")
	gomniauth.WithProviders(
		github.New("7ab63f7f8e4b9a02a0ee", "d0c2d4899ea03e7775ce7b6c3e7df3ed89f2b797", "http://localhost:8080/auth/callback/github"),
		google.New("831989856590-onu4to9s0aucvqt9pe3socp172ktfhu4.apps.googleusercontent.com", "dWr_140kcL2BtsevNlPS9Ay6", "http://localhost:8080/auth/callback/google"),
	)

	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()
	r := newRoom()

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	// start the web server
	http.ListenAndServe(*addr, nil)
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
	data := map[string]interface{}{
		"Hosta": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
		fmt.Printf("%v", data["UseraData"])
	}
	t.templ.Execute(w, data)
}

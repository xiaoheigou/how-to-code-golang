package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

type authHandler struct {
	next http.Handler
}

var (
	// You must register the app at https://github.com/settings/developers
	// Set callback to http://127.0.0.1:7000/github_oauth_cb
	// Set ClientId and ClientSecret to
	oauthConf = &oauth2.Config{
		ClientID:     "7ab63f7f8e4b9a02a0ee",
		ClientSecret: "d0c2d4899ea03e7775ce7b6c3e7df3ed89f2b797",
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     githuboauth.Endpoint,
	}
	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "thisshouldberandom"
)

// MustAuthreturn authhandler
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		//not authenticated
		fmt.Println("获取到cookie")
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		//some other error
		http.Error(w, err.Error(), http.StatusTemporaryRedirect)
		return
	}
	//success - call the next handler
	h.next.ServeHTTP(w, r)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		log.Println("TODO handle login for", provider)

		if err != nil {
			http.Error(w, fmt.Sprintf("err when type provider%s:%s", provider, err), http.StatusBadRequest)
			return
		}
		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("err when type beginAuth %s:%s", provider, err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		code := r.FormValue("code")
		fmt.Println(code)

		token, err := oauthConf.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		oauthClient := oauthConf.Client(oauth2.NoContext, token)
		client := github.NewClient(oauthClient)
		user, _, err := client.Users.Get(oauth2.NoContext, "")
		if err != nil {
			fmt.Printf("client.Users.Get() faled with '%s'\n", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)
		authCookieValue := objx.New(map[string]interface{}{
			"name": *user.Login,
		}).MustBase64()

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "auth", Value: authCookieValue, Expires: expiration}
		http.SetCookie(w, &cookie)

		fmt.Println(authCookieValue)

		w.Header().Set("Location", "/chat")
		fmt.Print("hhhhh")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

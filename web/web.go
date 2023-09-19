package web

import (
	"fmt"
	"net/http"
	"os"

	_ "example.com/m/auth"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var webRd *render.Render
var store = sessions.NewCookieStore([]byte(os.Getenv("UUID_SESSION_KEY")))

type Context struct {
	Text string `json:"text"`
}

func getSessionId(r *http.Request) (string, error) {
	session, err := store.Get(r, "session")
	fmt.Println("session", session)
	if err != nil {
		return "", err
	}
	val := session.Values["id"]
	if val == nil {
		return "", nil
	}
	return val.(string), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	sessionId, err := getSessionId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if sessionId == "" {
		http.Redirect(w, r, "/auth/google/login", http.StatusTemporaryRedirect)
	}
	webRd.HTML(w, http.StatusOK, "main", nil)
}

func WebHttpHandler() http.Handler {
	r := mux.NewRouter()
	webRd = render.New(render.Options{
		Directory:  "public",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "main",
	})
	n := negroni.Classic()

	n.UseHandler(r)
	r.HandleFunc("/", indexHandler)

	return n
}

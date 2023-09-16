package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Context struct {
	Text string `json:"text"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
}

func WebHttpHandler() http.Handler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.UseHandler(r)
	r.HandleFunc("/", indexHandler)

	return n
}

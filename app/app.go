package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type AppHandler struct {
	http.Handler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func MyHttpHandler() *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
	}

	r.HandleFunc("/", a.indexHandler)

	return a
}

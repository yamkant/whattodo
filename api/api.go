package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type APIHandler struct {
	http.Handler
}

func (a *APIHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Handler Index")
}

func APIHttpHandler() *APIHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(r)

	a := &APIHandler{
		Handler: n,
	}

	r.HandleFunc("/api/", a.indexHandler)

	return a
}

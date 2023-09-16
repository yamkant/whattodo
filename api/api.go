package api

import (
	"fmt"
	"net/http"

	"example.com/m/model"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()

// TODO: Modify schema format
type Success struct {
	Flag bool `json:"success"`
}

type APIHandler struct {
	http.Handler
}

func (a *APIHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Handler Index")
}

func (a *APIHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	todo := model.AddTodo(content)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *APIHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	todo := model.AddTodo(content)
	rd.JSON(w, http.StatusCreated, todo)
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

	r.HandleFunc("/api/todos/", a.addTodoHandler).Methods("POST")

	return a
}

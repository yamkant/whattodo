package api

import (
	"fmt"
	"net/http"
	"strconv"

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
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
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
	r.HandleFunc("/api/todos/", a.removeTodoHandler).Methods("DELETE")

	return a
}

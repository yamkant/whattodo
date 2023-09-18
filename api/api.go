package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/m/auth"
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

type AddTodoDTO struct {
	Content string
}

type UpdateTodoDTO struct {
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Completed bool      `json:"completed"`
}

type APIHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *APIHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API Handler Index")
}

func (a *APIHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := auth.GetSessionID(r)
	list := a.db.GetTodos(sessionId)
	rd.JSON(w, http.StatusOK, list)
}

func (a *APIHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	var body AddTodoDTO
	sessionId := auth.GetSessionID(r)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, nil)
	}
	todo := a.db.AddTodo(sessionId, body.Content)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *APIHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *APIHandler) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var body UpdateTodoDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, nil)
	}

	ok := a.db.UpdateTodo(id, body.Completed, body.StartedAt, body.EndedAt)
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
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/api/", a.indexHandler)

	r.HandleFunc("/api/todos/", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/api/todos/", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/api/todos/{id}/", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/api/todos/{id}/", a.updateTodoHandler).Methods("PATCH")

	return a
}

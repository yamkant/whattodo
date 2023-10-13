package handlers

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) apiStatus() {
	server.Router.GET("/api", Welcome)
}

func (server *Server) apiTodoRoutes() {
	server.Router.GET("/api/v1/todos/", GetTodos)
	server.Router.POST("/api/v1/todos/", AddTodo)
	server.Router.PATCH("/api/v1/todos/:id/", UpdateTodo)
}

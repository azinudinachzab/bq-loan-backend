package controllers

import (
	"github.com/azinudinachzab/bq-loan-backend/api/constant"
	"github.com/azinudinachzab/bq-loan-backend/api/middlewares"
	"github.com/azinudinachzab/bq-loan-backend/api/responses"
	"net/http"
)

func (server *Server) initializeRoutes() {

	// Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.healthCheck)).Methods(constant.HTTPMethodGet)

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods(constant.HTTPMethodPost)

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteUser)).Methods(constant.HTTPMethodDelete)

	//Posts routes
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.CreatePost)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.GetPosts)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(server.GetPost)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdatePost))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(server.DeletePost)).Methods(constant.HTTPMethodDelete)
}

func (server *Server) healthCheck(w http.ResponseWriter, _ *http.Request) {
	responses.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "OK"})
}

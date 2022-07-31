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
	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteUser)).Methods(constant.HTTPMethodDelete)
	server.Router.HandleFunc("/user/active/{id}", middlewares.SetMiddlewareAuthentication(server.ToggleUserIsActive)).Methods(constant.HTTPMethodPut)

	//Loan routes
	// Loan Type
	server.Router.HandleFunc("/loan-type", middlewares.SetMiddlewareJSON(server.CreateLoanType)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/loan-type", middlewares.SetMiddlewareJSON(server.GetLoanTypes)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(server.GetLoanType)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateLoanType))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteLoanType)).Methods(constant.HTTPMethodDelete)

	// Loan General
	server.Router.HandleFunc("/loan-general", middlewares.SetMiddlewareJSON(server.CreateLoanGeneral)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/loan-general", middlewares.SetMiddlewareJSON(server.GetLoanGenerals)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(server.GetLoanGeneral)).Methods(constant.HTTPMethodGet)
	//server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateLoanType))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteLoanGeneral)).Methods(constant.HTTPMethodDelete)
}

func (server *Server) healthCheck(w http.ResponseWriter, _ *http.Request) {
	responses.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "OK"})
}

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
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodOptions)

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods(constant.HTTPMethodOptions)

	//Users routes
	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.DeleteUser))).Methods(constant.HTTPMethodDelete)
	server.Router.HandleFunc("/user/active/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.ToggleUserIsActive))).Methods(constant.HTTPMethodPut)

	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/user/active/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.ToggleUserIsActive))).Methods(constant.HTTPMethodOptions)

	//Loan routes
	// Loan Type
	server.Router.HandleFunc("/loan-type", middlewares.SetMiddlewareJSON(server.CreateLoanType)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/loan-type", middlewares.SetMiddlewareJSON(server.GetLoanTypes)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(server.GetLoanType)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateLoanType))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.DeleteLoanType))).Methods(constant.HTTPMethodDelete)

	server.Router.HandleFunc("/loan-type", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/loan-type/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodOptions)

	// Loan General
	server.Router.HandleFunc("/loan-general", middlewares.SetMiddlewareJSON(server.CreateLoanGeneral)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/loan-general", middlewares.SetMiddlewareJSON(server.GetLoanGenerals)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(server.GetLoanGeneral)).Methods(constant.HTTPMethodGet)
	//server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateLoanType))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.DeleteLoanGeneral))).Methods(constant.HTTPMethodDelete)

	server.Router.HandleFunc("/loan-general", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/loan-general/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodOptions)

	// Loan Detail
	server.Router.HandleFunc("/loan-detail", middlewares.SetMiddlewareJSON(server.CreateLoanDetail)).Methods(constant.HTTPMethodPost)
	server.Router.HandleFunc("/loan-detail", middlewares.SetMiddlewareJSON(server.GetLoanDetails)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-detail/{id}", middlewares.SetMiddlewareJSON(server.GetLoanDetail)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-detail/general/{id}", middlewares.SetMiddlewareJSON(server.GetLoanDetailByGeneralID)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-detail/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateLoanDetail))).Methods(constant.HTTPMethodPut)
	server.Router.HandleFunc("/loan-detail/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.DeleteLoanDetail))).Methods(constant.HTTPMethodDelete)

	server.Router.HandleFunc("/loan-detail", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/loan-detail/general/{id}", middlewares.SetMiddlewareJSON(server.GetLoanDetailByGeneralID)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/loan-detail/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods(constant.HTTPMethodOptions)

	// Loan action
	server.Router.HandleFunc("/loan-general/accept/{id}", middlewares.SetMiddlewareJSON(server.AcceptLoanRequest)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-general/accept/{id}", middlewares.SetMiddlewareJSON(server.AcceptLoanRequest)).Methods(constant.HTTPMethodOptions)
	server.Router.HandleFunc("/loan-detail/accept/{id}", middlewares.SetMiddlewareJSON(server.AcceptLoanPayment)).Methods(constant.HTTPMethodGet)
	server.Router.HandleFunc("/loan-detail/accept/{id}", middlewares.SetMiddlewareJSON(server.AcceptLoanPayment)).Methods(constant.HTTPMethodOptions)
}

func (server *Server) healthCheck(w http.ResponseWriter, _ *http.Request) {
	responses.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "OK"})
}

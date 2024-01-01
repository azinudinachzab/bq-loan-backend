package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/azinudinachzab/bq-loan-be-v2/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

type HttpServer struct {
	service service.Service
}

func NewHttpServer(svc service.Service) http.Handler {
	r := chi.NewRouter()
	d := &HttpServer{
		service: svc,
	}

	/* ***** ***** *****
	 * init middleware
	 * ***** ***** *****/
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"*"}, // "True-Client-IP", "X-Forwarded-For", "X-Real-IP", "X-Request-Id", "Origin", "Accept", "Content-Type", "Authorization", "Token"
		AllowCredentials: true,
		MaxAge:           86400,
	}))
	r.Use(httprate.LimitByIP(80, 1*time.Minute))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	/* ***** ***** *****
	 * init custom error for 404 and 405
	 * ***** ***** *****/
	r.NotFound(d.Custom404)
	r.MethodNotAllowed(d.Custom405)

	/* ***** ***** *****
	 * init path route
	 * ***** ***** *****/
	r.Get("/", d.Home)
	r.Get("/dashboard", d.Dashboard)
	r.Get("/admin/dashboard", d.DashboardAdmin)

	// onboarding
	r.Post("/user", d.Registration)
	r.Post("/login", d.Login)
	r.Put("/user/active/{id}", d.ToggleIsActive)
	r.Post("/user/password", d.ChangePassword)

	// user
	r.Get("/user", d.GetUsers)
	r.Get("/user/{id}", d.GetUser)
	r.Put("/user/{id}", d.UpdateUser)
	r.Delete("/user/{id}", d.DeleteUser)

	// loan
	r.Post("/loan-general", d.CreateLoanGeneral)
	r.Put("/loan-general/{id}", d.UpdateLoanGeneral)
	r.Delete("/loan-general/{id}", d.DeleteLoanGeneral)
	r.Get("/loan-general", d.GetLoanGenerals)
	r.Get("/loan-general/{id}", d.GetLoanGeneral)

	//r.Put("/loan-detail/{id}", d.UpdateLoanDetail)
	r.Delete("/loan-detail/{lgid}/{id}", d.DeleteLoanDetail)
	r.Get("/loan-detail", d.GetLoanDetails)
	r.Get("/loan-detail/{id}", d.GetLoanDetail)

	r.Get("/loan-general/accept/{id}", d.AcceptLoanRequest)
	r.Get("/loan-detail/accept/{id}", d.AcceptPaymentRequest)

	// loan type
	r.Post("/loan-type", d.CreateLoanType)
	r.Put("/loan-type/{id}", d.UpdateLoanType)
	r.Delete("/loan-type/{id}", d.DeleteLoanType)
	r.Get("/loan-type", d.GetLoanTypes)
	r.Get("/loan-type/{id}", d.GetLoanType)

	return r
}

func (d *HttpServer) Home(w http.ResponseWriter, r *http.Request) {
	responseData(w, r, httpResponse{Message: "Hello World : " + time.Now().Format(time.RFC3339)})
}

func (d *HttpServer) AcceptLoanRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.AcceptLoanRequest(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Accept Loan Success",
	})
}

func (d *HttpServer) AcceptPaymentRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.AcceptPaymentRequest(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Accept Payment Success",
	})
}

func (d *HttpServer) Dashboard(w http.ResponseWriter, r *http.Request) {
	dashboard, err := d.service.Dashboard(r.Context())
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    dashboard,
	})
}

func (d *HttpServer) DashboardAdmin(w http.ResponseWriter, r *http.Request) {
	dashboard, err := d.service.DashboardAdmin(r.Context())
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    dashboard,
	})
}

func (d *HttpServer) Custom404(w http.ResponseWriter, r *http.Request) {
	err := errs.New(model.ECodeNotFound, "route does not exist")
	responseError(w, r, err)
}

func (d *HttpServer) Custom405(w http.ResponseWriter, r *http.Request) {
	err := errs.New(model.ECodeMethodFail, "method is not valid")
	responseError(w, r, err)
}

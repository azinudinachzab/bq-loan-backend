package delivery

import (
	"encoding/json"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (d *HttpServer) CreateLoanType(w http.ResponseWriter, r *http.Request) {
	var req model.LoanType

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	err := d.service.CreateLoanType(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusCreated,
		Message: "Registration Success",
	})
}

func (d *HttpServer) UpdateLoanType(w http.ResponseWriter, r *http.Request) {
	var req model.LoanType

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	req.ID = uint32(id)
	err = d.service.UpdateLoanType(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Update Success",
	})
}

func (d *HttpServer) DeleteLoanType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.DeleteLoanType(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	})
}

func (d *HttpServer) GetLoanType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	lt, err := d.service.GetLoanType(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    lt,
	})
}

func (d *HttpServer) GetLoanTypes(w http.ResponseWriter, r *http.Request) {
	lts, err := d.service.GetLoanTypes(r.Context())
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get List Success",
		Data:    lts,
	})
}

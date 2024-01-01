package delivery

import (
	"encoding/json"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (d *HttpServer) CreateLoanGeneral(w http.ResponseWriter, r *http.Request) {
	var req model.LoanGeneral

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	err := d.service.CreateLoanGeneral(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusCreated,
		Message: "Registration Success",
	})
}

func (d *HttpServer) UpdateLoanGeneral(w http.ResponseWriter, r *http.Request) {
	var req model.LoanGeneral

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
	err = d.service.UpdateLoanGeneral(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Update Success",
	})
}

func (d *HttpServer) DeleteLoanGeneral(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.DeleteLoanGeneral(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	})
}

func (d *HttpServer) GetLoanGeneral(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	lg, err := d.service.GetLoanGeneral(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    lg,
	})
}

func (d *HttpServer) GetLoanGenerals(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var idint uint32 = 0
	if id != "" {
		tmp, err := strconv.Atoi(id)
		if err != nil {
			responseError(w, r, errs.New(model.ECodeBadRequest, "uid must be number"))
			return
		}
		idint = uint32(tmp)
	}

	uid := r.URL.Query().Get("uid")
	var uidint uint32 = 0
	if uid != "" {
		tmp, err := strconv.Atoi(uid)
		if err != nil {
			responseError(w, r, errs.New(model.ECodeBadRequest, "uid must be number"))
			return
		}
		uidint = uint32(tmp)
	}

	title := r.URL.Query().Get("title")

	lgs, err := d.service.GetLoanGenerals(r.Context(), idint, uidint, title)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get List Success",
		Data:    lgs,
	})
}

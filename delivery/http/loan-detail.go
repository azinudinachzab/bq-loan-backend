package delivery

import (
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

//func (d *HttpServer) UpdateLoanDetail(w http.ResponseWriter, r *http.Request) {
//	var req model.LoanDetail
//
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
//		responseError(w, r, err)
//		return
//	}
//
//	id, err := strconv.Atoi(chi.URLParam(r, "id"))
//	if err != nil {
//		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
//		return
//	}
//
//	if id == 0 {
//		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
//		return
//	}
//
//	req.ID = uint32(id)
//	err = d.service.UpdateLoanDetail(r.Context(), req)
//	if err != nil {
//		responseError(w, r, err)
//		return
//	}
//
//	responseData(w, r, httpResponse{
//		Status:  http.StatusOK,
//		Message: "Update Success",
//	})
//}

func (d *HttpServer) DeleteLoanDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	lgid, err := strconv.Atoi(chi.URLParam(r, "lgid"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "lgid must be number"))
		return
	}

	if lgid == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "lgid cant be empty"))
		return
	}

	err = d.service.DeleteLoanDetail(r.Context(), uint32(id), uint32(lgid))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	})
}

func (d *HttpServer) GetLoanDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	ld, err := d.service.GetLoanDetail(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    ld,
	})
}

func (d *HttpServer) GetLoanDetails(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	lds, err := d.service.GetLoanDetails(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get List Success",
		Data:    lds,
	})
}

package delivery

import (
	"encoding/json"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (d *HttpServer) CreateSocialFund(w http.ResponseWriter, r *http.Request) {
	var req model.SocialFund

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	err := d.service.SocialFundRequest(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusCreated,
		Message: "Social fund request Success",
	})
}

func (d *HttpServer) AcceptSocialFundRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.AcceptSocialFundRequest(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Accept Fund Success",
	})
}

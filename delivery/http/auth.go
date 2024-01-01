package delivery

import (
	"encoding/json"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"net/http"
)

func (d *HttpServer) Registration(w http.ResponseWriter, r *http.Request) {
	var req model.RegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	err := d.service.Registration(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{Message: "Registration Success", Status: http.StatusCreated})
}

func (d *HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	usr, err := d.service.Login(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{Message: "Login Success", Status: http.StatusOK, Data: usr})
}

func (d *HttpServer) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID          uint32 `json:"id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = errs.New(model.ECodeBadRequest, "failed to decode request body")
		responseError(w, r, err)
		return
	}

	err := d.service.ChangePassword(r.Context(), req.ID, req.OldPassword, req.NewPassword)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{Message: "Login Success", Status: http.StatusOK})
}

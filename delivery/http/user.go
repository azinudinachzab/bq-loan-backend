package delivery

import (
	"encoding/json"
	"github.com/azinudinachzab/bq-loan-be-v2/model"
	"github.com/azinudinachzab/bq-loan-be-v2/pkg/errs"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (d *HttpServer) ToggleIsActive(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.ToggleIsActive(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Update Success",
	})
}

func (d *HttpServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req model.User

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
	err = d.service.UpdateUser(r.Context(), req)
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Update Success",
	})
}

func (d *HttpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	err = d.service.DeleteUser(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Delete Success",
	})
}

func (d *HttpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id must be number"))
		return
	}

	if id == 0 {
		responseError(w, r, errs.New(model.ECodeBadRequest, "id cant be empty"))
		return
	}

	usr, err := d.service.GetUser(r.Context(), uint32(id))
	if err != nil {
		responseError(w, r, err)
		return
	}

	responseData(w, r, httpResponse{
		Status:  http.StatusOK,
		Message: "Get Data Success",
		Data:    usr,
	})
}

func (d *HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]string)
	for _, val := range []string{"name", "email"} {
		tmp := r.URL.Query().Get(val)
		if tmp != "" {
			filter[val] = tmp
		}
	}

	lts, err := d.service.GetUsers(r.Context(), filter)
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

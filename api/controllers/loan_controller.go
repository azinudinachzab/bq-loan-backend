package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/azinudinachzab/bq-loan-backend/api/models"
	"github.com/azinudinachzab/bq-loan-backend/api/responses"
	"github.com/azinudinachzab/bq-loan-backend/api/utils/formaterror"
	"github.com/gorilla/mux"
)

// ============================= LOAN TYPE =======================================

func (server *Server) CreateLoanType(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanType := models.LoanType{}
	err = json.Unmarshal(body, &loanType)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanType.Prepare()
	//err = post.Validate()
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	typeCreated, err := loanType.SaveLoanType(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, typeCreated.ID))
	responses.JSON(w, http.StatusCreated, typeCreated)
}

func (server *Server) GetLoanTypes(w http.ResponseWriter, r *http.Request) {
	loanType := models.LoanType{}
	loanTypes, err := loanType.FindAllLoanTypes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanTypes)
}

func (server *Server) GetLoanType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	loanType := models.LoanType{}
	loanTypeData, err := loanType.FindLoanTypeByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanTypeData)
}

func (server *Server) UpdateLoanType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	loanTypeData := models.LoanType{}
	err = json.Unmarshal(body, &loanTypeData)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	loanTypeData.Prepare()
	//err = postUpdate.Validate()
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	postUpdated, err := loanTypeData.UpdateALoanType(server.DB, uint32(pid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeleteLoanType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the post exist
	loanType := models.LoanType{}
	if err := loanType.DeleteALoanType(server.DB, uint32(pid)); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}

// ============================= LOAN GENERAL =======================================

func (server *Server) CreateLoanGeneral(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanGeneral := models.LoanGeneral{}
	err = json.Unmarshal(body, &loanGeneral)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanGeneral.Prepare()
	//err = post.Validate()
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	typeCreated, err := loanGeneral.SaveLoanGeneral(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, typeCreated.ID))
	responses.JSON(w, http.StatusCreated, typeCreated)
}

func (server *Server) GetLoanGenerals(w http.ResponseWriter, r *http.Request) {
	loanGeneral := models.LoanGeneral{}
	loanGenerals, err := loanGeneral.FindAllLoanGenerals(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanGenerals)
}

func (server *Server) GetLoanGeneral(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	loanGeneral := models.LoanGeneral{}
	loanGeneralData, err := loanGeneral.FindLoanGeneralByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanGeneralData)
}

//func (server *Server) UpdateLoanGeneral(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	pid, err := strconv.ParseUint(vars["id"], 10, 64)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	// Start processing the request data
//	loanTypeData := models.LoanType{}
//	err = json.Unmarshal(body, &loanTypeData)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//
//	loanTypeData.Prepare()
//	//err = postUpdate.Validate()
//	//if err != nil {
//	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
//	//	return
//	//}
//	postUpdated, err := loanTypeData.UpdateALoanType(server.DB, uint32(pid))
//
//	if err != nil {
//		formattedError := formaterror.FormatError(err.Error())
//		responses.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	responses.JSON(w, http.StatusOK, postUpdated)
//}

func (server *Server) DeleteLoanGeneral(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the post exist
	loanGeneral := models.LoanGeneral{}
	if err := loanGeneral.DeleteALoanGeneral(server.DB, uint32(pid)); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
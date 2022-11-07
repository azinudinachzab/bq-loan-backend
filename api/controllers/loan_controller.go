package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	creditor := "%" + r.URL.Query().Get("creditor") + "%"
	loan := "%" + r.URL.Query().Get("loan") + "%"
	lastTime := r.URL.Query().Get("timestamp")

	if lastTime != "" {
		lastTime = strings.Replace(lastTime, " ", "+", 1)
		tParse, err := time.Parse(time.RFC3339, lastTime)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		lastTime = tParse.Format(time.RFC3339)
	}

	loanGeneral := models.LoanGeneral{}
	loanGenerals, err := loanGeneral.FindAllLoanGeneralsPaginatedSearch(creditor, loan, lastTime)
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

// ============================= LOAN DETAIL =======================================

func (server *Server) CreateLoanDetail(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanDetail := models.LoanDetail{}
	err = json.Unmarshal(body, &loanDetail)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	loanDetail.Prepare()
	//err = post.Validate()
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	typeCreated, err := loanDetail.SaveLoanDetail(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, typeCreated.ID))
	responses.JSON(w, http.StatusCreated, typeCreated)
}

func (server *Server) GetLoanDetails(w http.ResponseWriter, r *http.Request) {
	loanDetail := models.LoanDetail{}
	loanDetails, err := loanDetail.FindAllLoanDetails(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanDetails)
}

func (server *Server) GetLoanDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	loanDetail := models.LoanDetail{}
	loanDetailData, err := loanDetail.FindLoanDetailByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanDetailData)
}

func (server *Server) GetLoanDetailByGeneralID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	loanDetail := models.LoanDetail{}
	loanDetailData, err := loanDetail.FindLoanDetailByLoanGeneralID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, loanDetailData)
}

func (server *Server) UpdateLoanDetail(w http.ResponseWriter, r *http.Request) {

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
	loanDetail := models.LoanDetail{}
	err = json.Unmarshal(body, &loanDetail)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	loanDetail.Prepare()
	//err = postUpdate.Validate()
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	postUpdated, err := loanDetail.UpdateALoanDetail(server.DB, uint32(pid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeleteLoanDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the post exist
	loanDetail := models.LoanDetail{}
	if err := loanDetail.DeleteALoanDetail(server.DB, uint32(pid)); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}

// Loan action
func (server *Server) AcceptLoanRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	loanGeneral := models.LoanGeneral{}
	loanGeneralData, err := loanGeneral.FindLoanGeneralByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	tenor := loanGeneralData.Tenor
	loanDetails := make([]models.LoanDetail, 0)
	loanDetail := models.LoanDetail{}
	tmp := models.LoanDetail{
		LoanGeneralID: loanGeneralData.ID,
		Amount:        loanGeneralData.Amount / float64(tenor),
		Status:        0,
	}

	now := time.Now()
	for i:=0; i<tenor; i++{
		tmp.Prepare()
		date := time.Date(now.Year(), now.Month(), 25, now.Hour(), now.Minute(), now.Second(),
			now.Nanosecond(), now.Location()).AddDate(0, i+1, 0)
		tmp.Datetime = date

		loanDetails = append(loanDetails, tmp)
	}

	if len(loanDetails) > 0 {
		err = loanDetail.BulkSaveLoanDetail(server.DB, loanDetails)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
	}

	err = loanGeneral.UpdateStatus(server.DB, uint32(pid), 1)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) AcceptLoanPayment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	loanDetail := models.LoanDetail{}
	_, err = loanDetail.FindLoanDetailByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	err = loanDetail.UpdateStatus(server.DB, uint32(pid), 1)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, "")
}

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/azinudinachzab/bq-loan-backend/api/auth"
	"github.com/azinudinachzab/bq-loan-backend/api/models"
	"github.com/azinudinachzab/bq-loan-backend/api/responses"
	"github.com/azinudinachzab/bq-loan-backend/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	usr, token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, struct {
		Token string      `json:"token"`
		Usr   models.User `json:"user"`
	}{Token: token, Usr: usr})
}

func (server *Server) SignIn(email, password string) (models.User, string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return models.User{}, "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return models.User{}, "", err
	}
	tkn, err := auth.CreateToken(user.ID)
	return user, tkn, err
}

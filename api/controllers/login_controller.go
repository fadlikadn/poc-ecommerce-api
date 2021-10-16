package controllers

import (
	"encoding/json"
	"github.com/fadlikadn/poc-ecommerce-api/api/auth"
	"github.com/fadlikadn/poc-ecommerce-api/api/models"
	"github.com/fadlikadn/poc-ecommerce-api/api/responses"
	"github.com/fadlikadn/poc-ecommerce-api/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	customer := models.Customer{}

	err = server.DB.Debug().Model(models.Customer{}).Where("email = ?", email).Take(&customer).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(customer.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(uint32(customer.ID))
}


func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer.Prepare()
	err = customer.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(customer.Email, customer.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}


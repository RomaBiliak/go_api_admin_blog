package controllers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"golang.org/x/crypto/bcrypt"
	"github.com/RomaBiliak/go_api_admin_blog/api/models"
	"github.com/RomaBiliak/go_api_admin_blog/api/responses"
	"github.com/RomaBiliak/go_api_admin_blog/api/auth"
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

	token, err := server.SignIn(user.Login, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(login, password string) (string, error) {

	var err error

	user := models.User{}

	err = user.GetUserByLogin(server.db(), login)
	if err != nil {
		return "", err
	}
	
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.Id)
}
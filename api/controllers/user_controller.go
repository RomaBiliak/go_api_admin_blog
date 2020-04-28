package controllers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"reflect"
	"github.com/gorilla/mux"
	"github.com/RomaBiliak/go_api_admin_blog/api/models"
	"github.com/RomaBiliak/go_api_admin_blog/api/responses"
)


func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	filters, err := user.GetUsers(server.db())
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, filters)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.ParseInt(vars["id"], 10, 64)
	
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}

	err = user.GetUserById(server.db(), int64(user_id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id, err := strconv.ParseInt(vars["id"], 10, 64)
	
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}

	err = user.DeleteUserById(server.db(), int64(user_id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
func (server *Server) CRUDUser(w http.ResponseWriter, r *http.Request){

	user, err := getUser(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	inputs := make([]reflect.Value,1)
    inputs[0] = reflect.ValueOf(server.db())
	
	result := reflect.ValueOf(&user).MethodByName(server.CRUD[r.Method]+"User").Call(inputs)

	err, status := result[1].Interface().(error)
	if status {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	data := make(map[string]int64)
	data["id"] = result[0].Interface().(int64)
	responses.JSON(w, http.StatusOK, data)
}

func getUser(r *http.Request)(models.User, error){
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	return user, err
}

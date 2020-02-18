package controllers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/RomaBiliak/go_api_admin_blog/api/models"
	"github.com/RomaBiliak/go_api_admin_blog/api/responses"
	"reflect"
)


func (server *Server) GetFilters(w http.ResponseWriter, r *http.Request) {
	filter := models.Filter{}

	filters, err := filter.GetFilters(server.db())
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, filters)
}

func (server *Server) GetFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter_id, err := strconv.ParseInt(vars["id"], 10, 64)
	
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	filter := models.Filter{}

	err = filter.GetFilterById(server.db(), int64(filter_id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, filter)

}

func (server *Server) CRUDFilter(w http.ResponseWriter, r *http.Request){

	filter, err := getFilter(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	inputs := make([]reflect.Value,1)
    inputs[0] = reflect.ValueOf(server.db())
	
	result := reflect.ValueOf(&filter).MethodByName(server.CRUD[r.Method]+"Filter").Call(inputs)

	err, status := result[1].Interface().(error)
	if status {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	data := make(map[string]int64)
	data["id"] = result[0].Interface().(int64)
	responses.JSON(w, http.StatusOK, data)
}

func getFilter(r *http.Request)(models.Filter, error){
	filter := models.Filter{}
	err := json.NewDecoder(r.Body).Decode(&filter)
	return filter, err
}

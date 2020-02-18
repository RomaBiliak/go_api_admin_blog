package controllers

import (

	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/RomaBiliak/go_api_admin_blog/api/models"
	"github.com/RomaBiliak/go_api_admin_blog/api/responses"
	"reflect"
	"fmt"
)


func (server *Server) GetImages(w http.ResponseWriter, r *http.Request) {
	image := models.Image{}

	images, err := image.GetImages(server.db())
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, images)
}

func (server *Server) GetImage(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	image_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	image := models.Image{}

	err = image.GetImageById(server.db(), int64(image_id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, image)

}

func (server *Server) CRUDImage(w http.ResponseWriter, r *http.Request){

	image, err := getImage(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	inputs := make([]reflect.Value,1)
    inputs[0] = reflect.ValueOf(server.db())
	
	result := reflect.ValueOf(&image).MethodByName(server.CRUD[r.Method]+"Image").Call(inputs)

	err, status := result[1].Interface().(error)
	if status {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	data := make(map[string]int64)
	data["id"] = result[0].Interface().(int64)
	responses.JSON(w, http.StatusOK, data)
}

func getImage(r *http.Request)(models.Image, error){
	image := models.Image{}
	err := json.NewDecoder(r.Body).Decode(&image)
	fmt.Println(image)
	return image, err
}

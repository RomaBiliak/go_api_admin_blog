package controllers

import (
	"net/http"
	"errors"
	"io/ioutil"
	"strconv"
	"mime"
	"os"
	"github.com/RomaBiliak/go_api_admin_blog/api/responses"
	"github.com/RomaBiliak/go_api_admin_blog/library/image"
	"github.com/RomaBiliak/go_api_admin_blog/library/file"
)

func (server *Server) UploadImage(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.ParseInt(os.Getenv("MAX_UPLOAD_FILE_SIZE"), 10, 64)
	err := r.ParseMultipartForm(size)
	if err != nil {
        responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	
	image_file, handler, err := r.FormFile("image")

    if err != nil {
        responses.ERROR(w, http.StatusBadRequest, err)
		return
    }
    defer image_file.Close()
   
    fileBytes, err := ioutil.ReadAll(image_file)
    if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	//validate content type
	filetype := handler.Header["Content-Type"][0]
    if filetype != "image/jpeg" && filetype != "image/jpg" && filetype != "image/png"  {
		responses.ERROR(w, http.StatusBadRequest, errors.New("incorrect file format"))	
        return
    }
	//get end name file (format)
	file_format, err := mime.ExtensionsByType(handler.Header["Content-Type"][0])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	file.MkdirIfNotExist(os.Getenv("UPLOAD_IMAGE_DIR"))
    tempFile, err := ioutil.TempFile(os.Getenv("UPLOAD_IMAGE_DIR"), "*"+file_format[0])
    if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
    defer tempFile.Close()
    
    // write this byte array to our temporary file
	tempFile.Write(fileBytes)

	//AddWatermark
	name, err := image.AddWatermark(tempFile.Name())
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	//Resize
	_, err = image.Resize(name)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	//Split
	_, err = image.Split(name)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	data := make(map[string]string)
	data["path"] = file.GetName(tempFile.Name()) + ".jpg"
	responses.JSON(w, http.StatusOK, data)
}

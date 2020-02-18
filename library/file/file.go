package file

import (
	"strings"
	"os"
)

func GetName(path string)(string){
	name_arr := strings.Split(path, ".")
	name_arr = strings.Split(name_arr[0], "/")
	name_arr = strings.Split(name_arr[len(name_arr)-1], "\\")
	return  name_arr[len(name_arr)-1]
}

func MkdirIfNotExist(folder string)(error){
	_, err := os.Stat(folder); 
	if os.IsNotExist(err) {
		os.MkdirAll(folder, os.ModePerm)
	}
	return err
}
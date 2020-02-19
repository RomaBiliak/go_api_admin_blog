package controllers

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)
var CRUD = make(map[string]string)
type Server struct {
	Router *mux.Router
	CRUD	map[string]string
}
func (server *Server) Initialize() {
	server.CRUD = CRUD
	server.CRUD["POST"] = "Add"
	server.CRUD["PUT"] = "Edit"
	server.CRUD["PATCH"] = "Patch"
	server.CRUD["DELETE"] = "Delete"

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}
func (server *Server) db() *sql.DB{
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		fmt.Printf("Cannot connect to %s database")
		log.Fatal("This is the error:", err)
	}
	return db
}
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port "+addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
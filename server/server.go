package server
import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}
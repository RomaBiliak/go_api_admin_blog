package controllers

import (
	"github.com/RomaBiliak/go_api_admin_blog/api/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/login", s.Login).Methods("POST")

	s.Router.HandleFunc("/", middlewares.SetMiddlewareAuthentication(s.GetImages)).Methods("GET")
	s.Router.HandleFunc("/images", middlewares.SetMiddlewareAuthentication(s.GetImages)).Methods("GET")
	s.Router.HandleFunc("/image/{id}", middlewares.SetMiddlewareAuthentication(s.GetImage)).Methods("GET")
	s.Router.HandleFunc("/image", middlewares.SetMiddlewareAuthentication(s.CRUDImage)).Methods("POST")
	s.Router.HandleFunc("/image", middlewares.SetMiddlewareAuthentication(s.CRUDImage)).Methods("PUT")
	s.Router.HandleFunc("/image", middlewares.SetMiddlewareAuthentication(s.CRUDImage)).Methods("PATCH")
	s.Router.HandleFunc("/image", middlewares.SetMiddlewareAuthentication(s.CRUDImage)).Methods("DELETE")
	s.Router.HandleFunc("/upload_image", middlewares.SetMiddlewareAuthentication(s.UploadImage)).Methods("POST")

	s.Router.HandleFunc("/filters", middlewares.SetMiddlewareAuthentication(s.GetFilters)).Methods("GET")
	s.Router.HandleFunc("/filter/{id}", middlewares.SetMiddlewareAuthentication(s.GetFilter)).Methods("GET")
	s.Router.HandleFunc("/filter", middlewares.SetMiddlewareAuthentication(s.CRUDFilter)).Methods("POST")
	s.Router.HandleFunc("/filter", middlewares.SetMiddlewareAuthentication(s.CRUDFilter)).Methods("PUT")
	s.Router.HandleFunc("/filter", middlewares.SetMiddlewareAuthentication(s.CRUDFilter)).Methods("DELETE")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("PUT")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareAuthentication(s.CRUDUser)).Methods("DELETE")

	//s.Router.Use(middlewares.SetMiddlewareAuthentication)
	s.Router.Use(middlewares.SetMiddlewareJSON)
}
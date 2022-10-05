package controllers

import (
	"net/http"

	"github.com/mygoapp/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	s.Router.HandleFunc("/sendPost", middlewares.SetMiddlewareJSON(s.PostParams)).Methods(http.MethodPost)
	s.Router.HandleFunc("/sendGet", middlewares.SetMiddlewareJSON(s.GetParams)).Methods("GET")
}

package controllers

import (
   "fmt"
   "log"
   "net/http"

   "github.com/gorilla/mux"

)

type Server struct {
 Router	*mux.Router
}

func (server *Server) Initialize() {
  server.Router = mux.NewRouter()
  server.initializeRoutes()
}

func (server *Server) Run(addr string)  {
  log.Println("Listening to port ", addr)
  log.Fatal(http.ListenAndServe(addr, server.Router))
}


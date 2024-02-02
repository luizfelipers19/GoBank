package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/luizfelipers19/gobank/db/sqlc"
)

// Server struct serves http requests into our banking service

type Server struct {
	store  *db.Store
	router gin.Engine
}

// Creates a new http server and router
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts")

	server.router = *router
	return server

}

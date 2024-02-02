package api

import "github.com/gin-gonic/gin"

type createAccountRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {

}

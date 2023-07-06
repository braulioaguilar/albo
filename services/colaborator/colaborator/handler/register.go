package handler

import (
	"albo/colaborator"

	"github.com/gin-gonic/gin"
)

func MakeHandler(srv colaborator.Service, app *gin.Engine) {
	hdlr := NewHandler(srv)
	api := app.Group("/api")
	api.GET("/colaborators/:character", hdlr.Get)
}

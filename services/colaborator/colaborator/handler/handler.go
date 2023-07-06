package handler

import (
	"albo/colaborator"
	"albo/pkg/albohttp"
	"albo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Service colaborator.Service
}

func NewHandler(svr colaborator.Service) *handler {
	return &handler{
		Service: svr,
	}
}

func (h *handler) Get(ctx *gin.Context) {
	character, _ := ctx.Params.Get("character")

	if !utils.Contains(utils.GetHeros(), character) {
		ctx.JSON(http.StatusBadRequest, albohttp.Failure("Character not valid"))
		return
	}

	result, err := h.Service.Get(character)
	if err != nil {
		ctx.JSON(http.StatusNotFound, albohttp.Failure(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

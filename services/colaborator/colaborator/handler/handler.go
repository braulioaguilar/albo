package handler

import (
	"albo/colaborator"
	"albo/pkg/albohttp"
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
	heros := []string{"ironman", "capamerica"}
	character, _ := ctx.Params.Get("character")

	if !contains(heros, character) {
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

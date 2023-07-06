package handler

import (
	"albo/colaborator"
	"albo/domain"
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

	item, err := h.Service.Get(character)
	if err != nil {
		ctx.JSON(http.StatusNotFound, albohttp.Failure(err.Error()))
		return
	}

	lastsync := item.CreatedAt.Format("02-01-2006 15:04:05")
	data := domain.ColaboratorDTO{
		LastSync: lastsync,
		Editor:   item.Editor,
		Writer:   item.Writer,
		Colorist: item.Colorist,
	}

	ctx.JSON(http.StatusOK, data)
}

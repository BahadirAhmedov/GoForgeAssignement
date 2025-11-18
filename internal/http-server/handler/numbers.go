package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"


	"goforge/internal/domain/models"
	"goforge/internal/lib/logger/sl"
	"goforge/internal/transport/http/request"
	"goforge/internal/transport/http/response"
)

func (p *Numbers) NumberAdd(ctx *gin.Context) {
	const op = "handlers.NumberAdd"

	log := p.log.With(
		slog.String("op", op),
	)
	
	var request request.Number

	err := ctx.BindJSON(&request)
	if err != nil {
		log.Error("failed to decode request body", sl.Err(err))
		ctx.JSON(http.StatusBadRequest, response.Error(INVALID_REQUEST, "failed to decode request body"))
		return
	}
	log.Info("request body decoded", slog.Any("request", request))

	model := models.Number{
		Value: request.Value,
	}
	
	numbers, err := p.numbersProvider.SaveNumber(model)
	if err != nil {
		log.Error("failed to save number", sl.Err(err))
		ctx.JSON(http.StatusInternalServerError, response.Error(INTERNAL_ERROR, "failed to add team"))
		return
	}

	log.Info("number added")
	ctx.JSON(http.StatusOK, gin.H{"numbers": numbers})

}
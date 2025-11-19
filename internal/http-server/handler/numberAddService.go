package handler

import (
	"log/slog"

	"goforge/internal/domain/models"
	// "goforge/internal/transport/http/request"
)

const (
	INVALID_REQUEST = "INVALID_REQUEST"
	INTERNAL_ERROR  = "INTERNAL_ERROR"
)


type Numbers struct {
	Log *slog.Logger
	NumbersProvider NumbersProvider
}


type NumbersProvider interface {
	SaveNumber(models.Number) ([]int, error)
}


func New(
	log slog.Logger,
	numbersProvider NumbersProvider,
) *Numbers {
	return &Numbers{
		Log:                 &log,
		NumbersProvider: numbersProvider,
	}
}




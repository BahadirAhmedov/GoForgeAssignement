package app

import (
	"log/slog"
	"goforge/internal/http-server/handler"
	"goforge/internal/storage/postgres"

)

func New(
	log *slog.Logger, 
	host string, 
	port int, 
	user string, 
	password string, 
	dbname string,
) (*handler.Numbers) {
	storage, err := postgres.New(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}

	err = postgres.CreateTable(storage)
	if err != nil {
		panic(err)
	}
	SubscriptionHandlers := handler.New(*log, storage)
	return SubscriptionHandlers
}
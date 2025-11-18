package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"goforge/internal/app"
	"goforge/internal/config"
	"goforge/internal/lib/logger/sl"

)


const (
	envLocal = "local"
	envDev =  "dev"
	envProd = "prod"
)

func main() {
	
	router := gin.Default()
	cfg := config.MustLoad()


	log := setupLogger(cfg.Env)


	handlers := app.New(log, cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password, cfg.Db.Name)


	router.POST("/numbers", handlers.NumberAdd)


	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen: %s\\n", sl.Err(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown: %s\\n", sl.Err(err))
		os.Exit(1)	
	}

	<-ctx.Done()
	log.Info("imeout of 5 seconds.")
	log.Info("Server exiting")	

}

func setupLogger(env string) (*slog.Logger){
	var log *slog.Logger
	
	switch env{
		case envLocal:
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envDev:
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)

		case envProd:
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
		}  
	return log
}


package main

import (
	"context"
	app "github.com/atauov/image-converter"
	"github.com/atauov/image-converter/internal/config"
	"github.com/atauov/image-converter/internal/handler"
	"github.com/atauov/image-converter/internal/lib/logger/sl"
	"github.com/atauov/image-converter/internal/ports/redismq"
	"github.com/atauov/image-converter/internal/repository"
	"github.com/atauov/image-converter/internal/repository/postgres"
	"github.com/atauov/image-converter/internal/service"
	"github.com/joho/godotenv"
	lg "log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		lg.Print("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting image-converter", slog.String("env", cfg.Env))

	log.Debug("debug messages are enabled")

	db, err := postgres.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Error("failed to init postgres", sl.Err(err))
		os.Exit(1)
	}

	s3conn, err := service.S3Connection(&cfg.S3Server)
	if err != nil {
		log.Error("failed to init S3 connection", sl.Err(err))
	}

	redisClient := redismq.NewRedisClient(&cfg.RedisServer)
	defer redisClient.Close()

	rmq := redismq.NewRedisMQ(redisClient)
	_ = rmq
	//TODO change layers, add rmq to service or adapter

	repo := repository.NewRepository(db)
	services := service.NewService(repo, s3conn)
	handlers := handler.NewHandler(services, &cfg.HTTPServer)

	srv := new(app.Server)

	go func() {
		if err = srv.Run(&cfg.HTTPServer, handlers.InitRoutes()); err != nil {
			log.Error("failed to start http server", sl.Err(err))
			return
		}
	}()

	log.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server...")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Error("server shutdown failed", sl.Err(err))
		return
	}

	if err = repository.CloseRepository(db); err != nil {
		log.Error("failed to close repository", sl.Err(err))
		return
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
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

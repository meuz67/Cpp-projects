package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/internal/auth"
	"app/internal/config"
	"app/internal/database"
	"app/internal/handler"
	appmw "app/internal/middleware"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	ctx := context.Background()
	pool, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, pool); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	sessions := auth.NewSessionStore(cfg.SessionSecret)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(appmw.AttachUser(repository.NewUserRepository(pool), sessions))

	renderer, err := handler.NewTemplateRenderer("templates/*.html")
	if err != nil {
		log.Fatalf("templates: %v", err)
	}
	e.Renderer = renderer

	userRepo := repository.NewUserRepository(pool)
	taskRepo := repository.NewTaskRepository(pool)

	authHandler := handler.NewAuthHandler(userRepo, sessions)
	authHandler.RegisterPublic(e)

	protected := e.Group("", appmw.RequireAuth(sessions))
	authHandler.RegisterProtected(protected)

	taskHandler := handler.NewTaskHandler(taskRepo)
	taskHandler.Register(protected)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	go func() {
		log.Printf("server listening on http://localhost%s", addr)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
}

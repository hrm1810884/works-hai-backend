package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/service/humandrawing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// CORSミドルウェアを適用
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "x-api-key",
		},
	}))

	e.POST("/human-drawing", humandrawing.UploadHandler)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Server started at %v", cfg.Server.Dev)
	err = e.Start(cfg.Server.Dev)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

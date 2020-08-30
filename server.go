package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"pharmacy-backend/utils"
)

func main() {
	var (
		err    error
		port   string //server port from env
		secret string //secret for jwt sign
	)

	// -------------------------
	// Reading env variables
	// -------------------------

	// Read server port from env
	port, err = utils.ReadEnv("PORT")
	if err != nil {
		log.Fatalf("Fatal error setting server port: %v", err)
	}
	log.Printf("Server port set to: %v", port)

	secret, err = utils.ReadEnv("SECRET")
	if err != nil {
		log.Fatalf("Fatal error setting secret: %v", err)
	}
	log.Println("JWT Secret set")

	// -------------------------
	// Echo setup
	// -------------------------
	e := echo.New()

	// Middleware config
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// -------------------------
	// Routes definition
	// -------------------------
	e.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// -------------------------
	// Run server
	// -------------------------
	e.Logger.Fatal(e.Start(":" + port))
}

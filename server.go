package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"pharmacy-backend/handler"
	"pharmacy-backend/utils"
	"time"
)

func main() {
	var (
		err    error
		port   string  //server port from env
		secret string  //secret for jwt sign
		dbUrl  string  //database url
		conn   *sql.DB //database connection pool
	)

	log.Printf("Starting environment init...")
	initStartTime := time.Now()

	// -------------------------
	// Read env variables
	// -------------------------

	// Read server port from env
	port, err = utils.ReadEnv("PORT")
	if err != nil {
		log.Fatalf("Fatal error setting server port: %v", err)
	}
	log.Printf("Server port set to: %v", port)

	// Read secret from env
	secret, err = utils.ReadEnv("SECRET")
	if err != nil {
		log.Fatalf("Fatal error setting secret: %v", err)
	}
	log.Println("JWT Secret set")

	// Read db url
	dbUrl, err = utils.ReadEnv("DATABASE_URL")
	if err != nil {
		log.Fatalf("Fatal error setting database url: %v", err)
	}
	log.Println("DB URL set")

	// -------------------------
	// DB pool connection
	// -------------------------
	conn, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Can't connect to db: %v", err)
	}
	log.Printf("Connection string set")

	defer conn.Close()

	// Try ping db to check for availability
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Can't ping database: %v", err)
	}
	log.Printf("DB correctly pinged")

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
	e.POST("/login", handler.SignIn(conn, secret))

	// -------------------------
	// Run server
	// -------------------------
	initDuration := time.Since(initStartTime)
	log.Printf("Environment initialized in %s", initDuration)

	e.Logger.Fatal(e.Start(":" + port))
}

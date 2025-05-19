package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

	portStr := os.Getenv("PORT")
	env := os.Getenv("ENVIRONMENT")

	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		portInt = 4000
	}

	flag.IntVar(&cfg.port, "port", portInt, "api server port")
	flag.StringVar(&cfg.env, "env", env, "Environment (dev|staging|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}
	if cfg.env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.LoggerWithWriter(logger.Writer()), gin.Recovery())

	v1 := router.Group("/v1")
	v1.GET("/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("server failed: %v", err)
	}
}

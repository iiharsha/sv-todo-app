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
	"github.com/iiharsha/sv-todo-app/internal/jsonlog"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *jsonlog.Logger
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

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
	}
	if cfg.env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})
	err = srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}

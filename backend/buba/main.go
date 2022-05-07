package main

import (
	application "backend_movies/application"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var cfg application.Config

	flag.IntVar(&cfg.Port, "port", 4000, "default port")
	flag.StringVar(&cfg.Env, "env", "production", "default environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application.NewApplication(cfg, logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}

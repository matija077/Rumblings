package application

import "log"

type env = string
type Version = string

type Application struct {
	config Config
	logger *log.Logger
}

const version Version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type AppStatus struct {
	Status      string  `json:"status"`
	Environment env     `json:"environment"`
	Version     Version `json:"version"`
}

func NewApplication(cfg Config, logger *log.Logger) *Application {
	return &Application{
		config: cfg,
		logger: logger,
	}
}

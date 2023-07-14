package config

import (
	"log"
)

type Application struct {
	ErrorLog     *log.Logger
	infoLog      *log.Logger
	serverConfig *ServerConfig
}

// NewApplication returns a new Application struct
// func NewApplication(serverConfig *ServerConfig) *Application {}

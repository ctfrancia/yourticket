package config

import (
	"log"
)

// this file is where all configurations would be held related to the application

// Application is a struct that holds the application configuration
type Application struct {
	ErrorLog     *log.Logger
	infoLog      *log.Logger
	serverConfig *ServerConfig
}

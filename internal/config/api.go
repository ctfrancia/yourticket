package config

type ServerConfig struct {
	Addr  string
	Dsn   string
	Debug bool
}

// NewServerConfig returns a new ServerConfig struct
func NewServerConfig(addr string, dsn string, debug bool) *ServerConfig {
	return &ServerConfig{
		Addr:  addr,
		Dsn:   dsn,
		Debug: debug,
	}
}

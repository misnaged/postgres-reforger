package config

import "time"

// Scheme represents the application configuration scheme.
type Scheme struct {
	// Env is the application environment.
	Env      string
	Db       *Db
	Http     *Server
	Requests *Requests
}

type Requests struct {
	Ttl time.Duration
}

// Server represent basic server params
type Server struct {
	Port    int
	Timeout string
}

type Addr struct {
	Host string
	Port int
}

// Db is service Data base connection params
type Db struct {
	Addr     `mapstructure:",squash"`
	User     string
	Password string
	Name     string
	DbType   string
}

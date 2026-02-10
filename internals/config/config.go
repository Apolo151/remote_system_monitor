package config

import (
	"flag"
	"time"
)

type ServerConfig struct {
	Port     int
	Interval time.Duration
}

type ClientConfig struct {
	ServerAddr string
}

func LoadServerConfig() *ServerConfig {
	port := flag.Int("port", 50051, "Server port")
	interval := flag.Duration("interval", 5*time.Second, "Metrics collection interval")
	flag.Parse()

	return &ServerConfig{
		Port:     *port,
		Interval: *interval,
	}
}

func LoadClientConfig() *ClientConfig {
	serverAddr := flag.String("server", "localhost:50051", "Server address")
	flag.Parse()

	return &ClientConfig{
		ServerAddr: *serverAddr,
	}
}

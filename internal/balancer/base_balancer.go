package balancer

import (
	"log"
	"time"

	"github.com/franzinBr/sin-load-balancer/config"
	"github.com/franzinBr/sin-load-balancer/internal/server"
)

const (
	LeastConnections string = "least-connections"
	RoundRobin       string = "round-robin"
)

type Balancer interface {
	GetNextServer() (*server.Server, error)
}

type BaseBalancer struct {
	Servers []*server.Server
}


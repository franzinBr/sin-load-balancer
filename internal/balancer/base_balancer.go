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
	HealthCheck()
}

type BaseBalancer struct {
	Servers []*server.Server
}

func FactoryBalancer(config *config.Config, servers []*server.Server) Balancer {
	var balancer Balancer

	switch config.Sin.Method {
	case LeastConnections:
		balancer = NewLeastConnectionBalancer(servers)
	case RoundRobin:
		balancer = NewRoundRobinBalancer(servers)
	}

	return balancer
}

func (b *BaseBalancer) HealthCheck() {

	config := config.GetConfig()

	interval, err := time.ParseDuration(config.Server.HealthCheck.CheckInterval)
	if err != nil {
		log.Fatalf("Invalid health check interval: %s", err.Error())
	}

	for _, server := range b.Servers {
		go server.CheckHealthy(
			config.Server.HealthCheck.Endpoint,
			config.Server.HealthCheck.StatusCode,
			interval,
		)
	}
}

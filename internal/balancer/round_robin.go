package balancer

import (
	"fmt"
	"sync"

	"github.com/franzinBr/sin-load-balancer/internal/server"
)

type RoundRobinBalancer struct {
	BaseBalancer
	Current int
	Mutex   sync.Mutex
}

func NewRoundRobinBalancer(servers []*server.Server) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		BaseBalancer: BaseBalancer{
			Servers: servers,
		},
		Current: -1,
	}
}

func (b *RoundRobinBalancer) rotate() *server.Server {
	b.Mutex.Lock()
	b.Current = (b.Current + 1) % len(b.Servers)
	b.Mutex.Unlock()
	return b.Servers[b.Current]
}

func (b *RoundRobinBalancer) GetNextServer() (*server.Server, error) {
	for i := 0; i < len(b.Servers); i++ {
		nextServer := b.rotate()

		if nextServer.Healthy {
			return nextServer, nil
		}
	}
	return nil, fmt.Errorf("Cannot found healthy server")
}

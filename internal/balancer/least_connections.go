package balancer

import "github.com/franzinBr/sin-load-balancer/internal/server"

type LeastConnectionBalancer struct {
	BaseBalancer
}

func NewLeastConnectionBalancer(servers []*server.Server) *LeastConnectionBalancer {
	return &LeastConnectionBalancer{
		BaseBalancer: BaseBalancer{
			Servers: servers,
		},
	}
}

func (b *LeastConnectionBalancer) GetNextServer() (*server.Server, error) {
	leastActiveConnections := -1
	leastActiveServer := b.Servers[0]
	for _, server := range b.Servers {
		server.Mutex.Lock()
		if (server.Connections < leastActiveConnections || leastActiveConnections == -1) && server.Healthy {
			leastActiveConnections = server.Connections
			leastActiveServer = server
		}
		server.Mutex.Unlock()
	}

	return leastActiveServer, nil

}

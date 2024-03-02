package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/franzinBr/sin-load-balancer/config"
)

type Server struct {
	URL         *url.URL
	Connections int
	Healthy     bool
	Mutex       sync.Mutex
}

func BuildServers(config *config.Config) []*Server {
	var servers []*Server
	for _, serverUrl := range config.Server.URLs {
		u, _ := url.Parse(serverUrl)
		servers = append(servers, &Server{URL: u, Healthy: true})
	}

	return servers
}

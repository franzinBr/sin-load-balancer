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

func (s *Server) CheckHealthy(endpoint string, sucessStatusCode int, interval time.Duration) {

	for range time.Tick(time.Duration(interval)) {
		url := s.URL.String() + endpoint

		res, err := http.Get(url)

		s.Healthy = err == nil && res != nil && res.StatusCode == sucessStatusCode
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mutex.Lock()
	s.Connections++
	s.Mutex.Unlock()

	proxy := httputil.NewSingleHostReverseProxy(s.URL)
	proxy.ServeHTTP(w, r)

	s.Mutex.Lock()
	s.Connections--
	s.Mutex.Unlock()
}

func BuildServers(config *config.Config) []*Server {
	var servers []*Server
	for _, serverUrl := range config.Server.URLs {
		u, _ := url.Parse(serverUrl)
		servers = append(servers, &Server{URL: u})
	}

	return servers
}

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

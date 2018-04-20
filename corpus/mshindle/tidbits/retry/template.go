package retry

import (
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

// RestTemplate provides a generic interface for multiple client packages
type RestTemplate interface {
	Get(url string) (resp *http.Response, err error)
}

/* --- standard http client --- */

type standard struct {
	host string
}

func (s *standard) Get(url string) (resp *http.Response, err error) {
	absURL := s.host + url
	return http.Get(absURL)
}

/* --- circuit breaker http client --- */

type breaker struct {
	scheme string
	hosts  []string
}

func (b *breaker) RequestHook(l *log.Logger, req *http.Request, retry int) {
	l.Printf("working on try: %d", retry)
	req.URL.Scheme = b.scheme
	req.URL.Host = b.hosts[retry]
}

func NewClient(hosts ...string) RestTemplate {
	if len(hosts) == 1 {
		return &standard{host: stripSlash(hosts[0])}
	}
	b := &breaker{scheme: "http", hosts: hosts}
	client := retryablehttp.NewClient()
	client.RequestLogHook = b.RequestHook
	client.RetryMax = len(b.hosts) - 1
	return client
}

func stripSlash(s string) string {
	return strings.TrimSuffix(s, "/")
}

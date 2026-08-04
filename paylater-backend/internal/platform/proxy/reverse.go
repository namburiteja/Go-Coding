package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Forward returns a Gin handler that reverse-proxies requests to targetBaseURL.
// Paths and headers (including Authorization) are preserved so JWT stays compatible.
func Forward(targetBaseURL string) gin.HandlerFunc {
	remote, err := url.Parse(targetBaseURL)
	if err != nil {
		panic("invalid proxy target URL: " + targetBaseURL)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = remote.Host
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

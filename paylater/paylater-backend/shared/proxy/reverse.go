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
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		http.Error(w, e.Error(), http.StatusBadGateway)
	}
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = remote.Host
		req.URL.Host = remote.Host
		req.URL.Scheme = remote.Scheme
	}

	return func(c *gin.Context) {
		// Prevent Gin from writing a 404/status after the proxied response.
		c.Abort()
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

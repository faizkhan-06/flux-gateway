package proxy

import (
	"log/slog"
	"net/http/httputil"
	"net/url"
)

func New(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		slog.Error("Invalid target URL provided to proxy", slog.Any("error", err))
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy
}

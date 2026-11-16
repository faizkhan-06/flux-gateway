package proxy

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/faizkhan-06/flux-gateway/internal/config"
)

func New(cfg *config.Gateway) *httputil.ReverseProxy {

	return &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			for _, route := range cfg.Routes {
				if strings.HasPrefix(pr.In.URL.Path, route.Path) {
					pr.SetURL(route.ParsedURL)
					pr.Out.Host = route.ParsedURL.Host

					slog.Info("Routing request",
						slog.String("path", pr.In.URL.Path),
						slog.String("routed_to", route.ParsedURL.String()),
					)
					return
				}
			}
			slog.Warn("No routing rule found for path", slog.String("path", pr.In.URL.Path))
		},

		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			if strings.Contains(err.Error(), "unsupported protocol scheme") {
				http.Error(w, "404 - Route Not Found in Gateway", http.StatusNotFound)
				return
			}
			slog.Error("Backend connection failed",
				slog.String("path", r.URL.Path),
				slog.Any("error", err),
			)
			http.Error(w, "502 - Bad Gateway", http.StatusBadGateway)
		},
	}
}

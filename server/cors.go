package server

import (
	"net/http"
	"strings"
)

func allowCORS(h http.Handler, allowOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			if len(allowOrigins) > 0 {
				for _, allowOrigin := range allowOrigins {
					if origin == allowOrigin {
						w.Header().Set("Access-Control-Allow-Origin", origin)
						break
					}
				}
			} else {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization", "Origin"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

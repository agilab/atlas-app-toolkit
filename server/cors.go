package server

import (
	"net/http"
	"strings"
)

var (
	defaultAllowHeaders = []string{"Content-Type", "Accept", "Authorization", "Origin"}
)

func allowCORS(h http.Handler, allowOrigins []string, extraAllowHeaders []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = r.URL.Scheme + "//" + r.URL.Host
		}

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
			preflightHandler(w, r, extraAllowHeaders)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func evaluateExtraAllowHeaders(allowHeaders []string) []string {
	m := map[string]bool{}
	for _, h := range defaultAllowHeaders {
		m[h] = true
	}

	extraAllowHeaders := []string{}
	for _, h := range allowHeaders {
		if m[h] == false {
			extraAllowHeaders = append(extraAllowHeaders, h)
		}
	}
	return extraAllowHeaders
}

func preflightHandler(w http.ResponseWriter, r *http.Request, extraAllowHeaders []string) {
	headers := defaultAllowHeaders
	if len(extraAllowHeaders) > 0 {
		headers = append(headers, extraAllowHeaders...)
	}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

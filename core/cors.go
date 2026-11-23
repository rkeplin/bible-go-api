package core

import (
	"net/http"
	"os"
	"strings"
)

// DefaultCORSOrigins are the browser origins allowed to call the API.
// Override with CORS_ALLOW_ORIGINS (comma separated).
var DefaultCORSOrigins = []string{
	"https://bible-ui.rkeplin.com",
	"https://k8s-bible-ui.rkeplin.com",
}

// CORSOriginsFromEnv returns CORS_ALLOW_ORIGINS split on commas, or
// DefaultCORSOrigins when it is unset or empty.
func CORSOriginsFromEnv() []string {
	var origins []string
	for _, o := range strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",") {
		if o = strings.TrimSpace(o); o != "" {
			origins = append(origins, o)
		}
	}
	if len(origins) == 0 {
		return DefaultCORSOrigins
	}
	return origins
}

// CORS reproduces exactly what the ingress-nginx enable-cors annotations used
// to do in front of this API, so moving CORS into the app changes nothing for
// clients:
//
//   - An Origin matches if it ends with an allowed origin, case-insensitively
//     (nginx used the regex `(origin1|origin2)$` with ~*). The Origin is echoed
//     back exactly as sent.
//   - A matching request gets all six CORS headers on every response,
//     whatever the status.
//   - A matching OPTIONS request is answered 204 right here, with or without
//     Access-Control-Request-Method, and never reaches the router.
//   - Anything else passes through untouched: no CORS headers, and no Vary.
func CORS(next http.Handler, allowedOrigins []string) http.Handler {
	allowed := make([]string, len(allowedOrigins))
	for i, o := range allowedOrigins {
		allowed[i] = strings.ToLower(o)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" || !originAllowed(strings.ToLower(origin), allowed) {
			next.ServeHTTP(w, r)
			return
		}

		h := w.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Credentials", "true")
		h.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cache-Control")
		h.Set("Access-Control-Expose-Headers", "Content-Length,Content-Range")
		h.Set("Access-Control-Max-Age", "1728000")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func originAllowed(origin string, allowed []string) bool {
	for _, a := range allowed {
		if strings.HasSuffix(origin, a) {
			return true
		}
	}
	return false
}

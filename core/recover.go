package core

import (
	"log"
	"net/http"
	"runtime/debug"
)

// Recover turns a handler panic into a 502 response instead of a dropped
// connection.
//
// A dropped connection used to reach clients as ingress-nginx's 502 (with
// the CORS headers still added). Envoy reports it as a 503 without them,
// so answering 502 here keeps what clients see the same behind either
// proxy. Panics are still logged with a stack trace, like net/http does.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			if err == http.ErrAbortHandler {
				panic(err)
			}
			log.Printf("http: panic serving %s: %v\n%s", r.RemoteAddr, err, debug.Stack())
			http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		}()
		next.ServeHTTP(w, r)
	})
}

package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// corsTestRouter mirrors the real router's behavior (StrictSlash, GET-only
// routes, gorilla's default 404/405) without needing a database.
func corsTestRouter() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods("GET").Path("/translations").HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte(`[]`))
	})
	return CORS(r, DefaultCORSOrigins)
}

var corsHeaders = map[string]string{
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Methods":     "GET, OPTIONS",
	"Access-Control-Allow-Headers":     "Origin, X-Requested-With, Content-Type, Accept, Cache-Control",
	"Access-Control-Expose-Headers":    "Content-Length,Content-Range",
	"Access-Control-Max-Age":           "1728000",
}

// Each case is taken from the live ingress-nginx responses captured before
// CORS moved into the app (cluster-infra test/cors-check.sh).
func TestCORSMatchesIngressNginx(t *testing.T) {
	const ui, k8s = "https://bible-ui.rkeplin.com", "https://k8s-bible-ui.rkeplin.com"

	cases := []struct {
		name, method, path, origin, acrm string
		wantStatus                       int
		wantCORS                         bool
	}{
		{"get ui", "GET", "/translations", ui, "", 200, true},
		{"get k8s ui", "GET", "/translations", k8s, "", 200, true},
		{"get other origin", "GET", "/translations", "https://evil.example", "", 200, false},
		{"get no origin", "GET", "/translations", "", "", 200, false},
		{"origin case-insensitive", "GET", "/translations", "HTTPS://BIBLE-UI.RKEPLIN.COM", "", 200, true},
		{"origin trailing slash", "GET", "/translations", ui + "/", "", 200, false},
		{"origin explicit port", "GET", "/translations", ui + ":443", "", 200, false},
		{"origin as subdomain prefix", "GET", "/translations", ui + ".evil.com", "", 200, false},
		{"origin with junk before (nginx regex is end-anchored only)", "GET", "/translations", "x" + ui, "", 200, true},
		{"origin http scheme", "GET", "/translations", "http://bible-ui.rkeplin.com", "", 200, false},
		{"origin null", "GET", "/translations", "null", "", 200, false},
		{"strict slash redirect", "GET", "/translations/", ui, "", 301, true},
		{"not found", "GET", "/nope-xyz", ui, "", 404, true},
		{"not found no origin", "GET", "/nope-xyz", "", "", 404, false},
		{"root", "GET", "/", ui, "", 404, true},
		{"preflight", "OPTIONS", "/translations", ui, "GET", 204, true},
		{"options without request-method", "OPTIONS", "/translations", k8s, "", 204, true},
		{"preflight unknown path", "OPTIONS", "/nope-xyz", ui, "GET", 204, true},
		{"preflight other origin", "OPTIONS", "/translations", "https://evil.example", "GET", 405, false},
		{"options no origin", "OPTIONS", "/translations", "", "", 405, false},
		{"head", "HEAD", "/translations", ui, "", 405, true},
		{"post", "POST", "/translations", ui, "", 405, true},
		{"post no origin", "POST", "/translations", "", "", 405, false},
	}

	h := corsTestRouter()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(c.method, c.path, nil)
			if c.origin != "" {
				req.Header.Set("Origin", c.origin)
			}
			if c.acrm != "" {
				req.Header.Set("Access-Control-Request-Method", c.acrm)
			}
			w := httptest.NewRecorder()
			h.ServeHTTP(w, req)

			if w.Code != c.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, c.wantStatus)
			}
			if v := w.Header().Values("Vary"); len(v) != 0 {
				t.Errorf("Vary = %q, want none", v)
			}
			if !c.wantCORS {
				for k := range w.Header() {
					if len(k) > 15 && k[:15] == "Access-Control-" {
						t.Errorf("unexpected header %s", k)
					}
				}
				return
			}
			if got := w.Header().Values("Access-Control-Allow-Origin"); len(got) != 1 || got[0] != c.origin {
				t.Errorf("Access-Control-Allow-Origin = %q, want [%q]", got, c.origin)
			}
			for k, want := range corsHeaders {
				if got := w.Header().Values(k); len(got) != 1 || got[0] != want {
					t.Errorf("%s = %q, want [%q]", k, got, want)
				}
			}
			if c.wantStatus == 204 && w.Body.Len() != 0 {
				t.Errorf("204 body = %q, want empty", w.Body.String())
			}
		})
	}
}

func TestCORSOriginsFromEnv(t *testing.T) {
	t.Setenv("CORS_ALLOW_ORIGINS", "")
	if got := CORSOriginsFromEnv(); len(got) != 2 || got[0] != DefaultCORSOrigins[0] {
		t.Errorf("unset: got %q, want defaults", got)
	}
	t.Setenv("CORS_ALLOW_ORIGINS", " https://a.example , ,https://b.example")
	if got := CORSOriginsFromEnv(); len(got) != 2 || got[0] != "https://a.example" || got[1] != "https://b.example" {
		t.Errorf("set: got %q", got)
	}
}

package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogger(t *testing.T) {
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := Logger(inner, "TestRoute")

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestNewRouter(t *testing.T) {
	routes.Add(Route{
		Name:        "RouterTest",
		Method:      "GET",
		Pattern:     "/router-test",
		HandlerFunc: http.NotFound,
	})

	router := NewRouter()
	if router == nil {
		t.Fatal("expected non-nil router")
	}

	req := httptest.NewRequest("GET", "/router-test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

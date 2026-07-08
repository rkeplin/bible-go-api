package core

import (
	"net/http"
	"testing"
)

func TestRoutesAddAndGetAll(t *testing.T) {
	r := &Routes{}

	if len(r.GetAll()) != 0 {
		t.Fatal("expected empty routes slice")
	}

	r.Add(Route{Name: "Test", Method: "GET", Pattern: "/test", HandlerFunc: http.NotFound})
	r.Add(Route{Name: "Test2", Method: "POST", Pattern: "/test2", HandlerFunc: http.NotFound})

	all := r.GetAll()
	if len(all) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(all))
	}
	if all[0].Name != "Test" {
		t.Errorf("expected route name 'Test', got %q", all[0].Name)
	}
	if all[1].Pattern != "/test2" {
		t.Errorf("expected pattern '/test2', got %q", all[1].Pattern)
	}
}

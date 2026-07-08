package core

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespond(t *testing.T) {
	payload := map[string]string{"key": "value"}

	w := httptest.NewRecorder()
	Respond(w, http.StatusOK, payload)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "application/json; charset=UTF-8" {
		t.Errorf("unexpected Content-Type: %q", ct)
	}

	var got map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if got["key"] != "value" {
		t.Errorf("expected key=value, got %v", got)
	}
}

func TestRespondNotFound(t *testing.T) {
	payload := HttpErrorResponse{Code: 404, Status: "Not Found", Message: "Item was not found."}

	w := httptest.NewRecorder()
	Respond(w, http.StatusNotFound, payload)

	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}

	var got HttpErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if got.Code != 404 {
		t.Errorf("expected code 404, got %d", got.Code)
	}
	if got.Status != "Not Found" {
		t.Errorf("expected status 'Not Found', got %q", got.Status)
	}
}

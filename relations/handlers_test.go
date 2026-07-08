package relations

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockRepo struct {
	collection []TextCollection
	err        error
}

func (m mockRepo) FindAll(_ int, _ string) ([]TextCollection, error) {
	return m.collection, m.err
}

func TestFindAll_OK(t *testing.T) {
	collection := []TextCollection{
		{
			{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "For God so loved...", Book: Book{ID: 43, Name: "John", Testament: "NT"}},
		},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/verse/1001001/relations", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1001001"})
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got []TextCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 1 {
		t.Errorf("expected 1 group, got %d", len(got))
	}
}

func TestFindAll_Empty(t *testing.T) {
	h := Handler{repo: mockRepo{collection: []TextCollection{}}}

	req := httptest.NewRequest("GET", "/verse/9999/relations", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "9999"})
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestFindAll_RepoError(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("db error")}}

	req := httptest.NewRequest("GET", "/verse/1/relations", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on repo error, got none")
		}
	}()
	h.FindAll(w, req)
}

func TestFindAll_WithTranslation(t *testing.T) {
	collection := []TextCollection{
		{
			{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "For God so loved..."},
		},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/verse/1001001/relations?translation=NIV", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1001001"})
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

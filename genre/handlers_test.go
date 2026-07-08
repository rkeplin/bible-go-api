package genre

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockRepo struct {
	genres GenreCollection
	genre  Genre
	err    error
}

func (m mockRepo) FindAll() (GenreCollection, error) { return m.genres, m.err }
func (m mockRepo) FindOne(_ int) (Genre, error)      { return m.genre, m.err }

func TestFindAll_OK(t *testing.T) {
	genres := GenreCollection{
		{ID: 1, Name: "Law"},
		{ID: 2, Name: "History"},
	}
	h := Handler{repo: mockRepo{genres: genres}}

	req := httptest.NewRequest("GET", "/genres", nil)
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got GenreCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 genres, got %d", len(got))
	}
	if got[0].Name != "Law" {
		t.Errorf("expected first genre 'Law', got %q", got[0].Name)
	}
}

func TestFindOne_OK(t *testing.T) {
	h := Handler{repo: mockRepo{genre: Genre{ID: 3, Name: "Poetry"}}}

	req := httptest.NewRequest("GET", "/genres/3", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "3"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got Genre
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got.Name != "Poetry" {
		t.Errorf("expected genre 'Poetry', got %q", got.Name)
	}
}

func TestFindOne_NotFound(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("not found")}}

	req := httptest.NewRequest("GET", "/genres/99", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "99"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestFindAll_RepoError(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("db error")}}

	req := httptest.NewRequest("GET", "/genres", nil)
	w := httptest.NewRecorder()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on repo error, got none")
		}
	}()
	h.FindAll(w, req)
}

package book

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockRepo struct {
	books    BookCollection
	book     Book
	chapters ChapterCollection
	err      error
}

func (m mockRepo) FindAll() (BookCollection, error)                              { return m.books, m.err }
func (m mockRepo) FindOne(_ int) (Book, error)                                   { return m.book, m.err }
func (m mockRepo) FindChapters(_ int, _ string) (ChapterCollection, error)       { return m.chapters, m.err }

func TestFindAll_OK(t *testing.T) {
	books := BookCollection{
		{ID: 1, Name: "Genesis", Testament: "OT"},
		{ID: 2, Name: "Exodus", Testament: "OT"},
	}
	h := Handler{repo: mockRepo{books: books}}

	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got BookCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 books, got %d", len(got))
	}
	if got[0].Name != "Genesis" {
		t.Errorf("expected first book Genesis, got %q", got[0].Name)
	}
}

func TestFindOne_OK(t *testing.T) {
	book := Book{ID: 1, Name: "Genesis", Testament: "OT", Genre: Genre{ID: 1, Name: "Law"}}
	h := Handler{repo: mockRepo{book: book}}

	req := httptest.NewRequest("GET", "/books/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got Book
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got.ID != 1 || got.Name != "Genesis" {
		t.Errorf("unexpected book: %+v", got)
	}
}

func TestFindOne_NotFound(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("not found")}}

	req := httptest.NewRequest("GET", "/books/999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "999"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestFindChapters_OK(t *testing.T) {
	chapters := ChapterCollection{{ID: 1}, {ID: 2}, {ID: 3}}
	h := Handler{repo: mockRepo{chapters: chapters}}

	req := httptest.NewRequest("GET", "/books/1/chapters", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	h.FindChapters(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got ChapterCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 3 {
		t.Errorf("expected 3 chapters, got %d", len(got))
	}
}

func TestFindChapters_NotFound(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("not found")}}

	req := httptest.NewRequest("GET", "/books/999/chapters", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "999"})
	w := httptest.NewRecorder()
	h.FindChapters(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

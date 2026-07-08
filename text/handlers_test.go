package text

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockRepo struct {
	collection TextCollection
	item       Text
	err        error
}

func (m mockRepo) FindAll(_ int, _ int, _ string) (TextCollection, error) {
	return m.collection, m.err
}
func (m mockRepo) FindOne(_ int, _ string) (Text, error) { return m.item, m.err }

func TestFindAllFromChapter_OK(t *testing.T) {
	collection := TextCollection{
		{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "In the beginning..."},
		{ID: 1001002, ChapterID: 1, VerseID: 2, Verse: "And the earth..."},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/books/1/chapters/1", nil)
	req = mux.SetURLVars(req, map[string]string{"bookId": "1", "chapterId": "1"})
	w := httptest.NewRecorder()
	h.FindAllFromChapter(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got TextCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 verses, got %d", len(got))
	}
}

func TestFindOne_OK(t *testing.T) {
	item := Text{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "In the beginning..."}
	h := Handler{repo: mockRepo{item: item}}

	req := httptest.NewRequest("GET", "/books/1/chapters/1/1001001", nil)
	req = mux.SetURLVars(req, map[string]string{"bookId": "1", "chapterId": "1", "verseId": "1001001"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got Text
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got.ID != 1001001 {
		t.Errorf("expected ID 1001001, got %d", got.ID)
	}
}

func TestFindOne_NotFound(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("not found")}}

	req := httptest.NewRequest("GET", "/books/1/chapters/1/9999999", nil)
	req = mux.SetURLVars(req, map[string]string{"bookId": "1", "chapterId": "1", "verseId": "9999999"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestFindAllFromChapter_WithTranslation(t *testing.T) {
	collection := TextCollection{
		{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "In the beginning..."},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/books/1/chapters/1?translation=ASV", nil)
	req = mux.SetURLVars(req, map[string]string{"bookId": "1", "chapterId": "1"})
	w := httptest.NewRecorder()
	h.FindAllFromChapter(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

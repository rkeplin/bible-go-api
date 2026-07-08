package translation

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockRepo struct {
	collection TranslationCollection
	item       Translation
	err        error
}

func (m mockRepo) FindAll() (TranslationCollection, error) { return m.collection, m.err }
func (m mockRepo) FindOne(_ int) (Translation, error)      { return m.item, m.err }

func TestFindAll_OK(t *testing.T) {
	collection := TranslationCollection{
		{ID: 1, Table: "t_kjv", Language: "English", Abbreviation: "KJV", Version: "King James Version"},
		{ID: 2, Table: "t_asv", Language: "English", Abbreviation: "ASV", Version: "American Standard Version"},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/translations", nil)
	w := httptest.NewRecorder()
	h.FindAll(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got TranslationCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 translations, got %d", len(got))
	}
}

func TestFindOne_OK(t *testing.T) {
	item := Translation{ID: 1, Table: "t_kjv", Language: "English", Abbreviation: "KJV", Version: "King James Version"}
	h := Handler{repo: mockRepo{item: item}}

	req := httptest.NewRequest("GET", "/translations/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got Translation
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got.Abbreviation != "KJV" {
		t.Errorf("expected KJV, got %q", got.Abbreviation)
	}
}

func TestFindOne_NotFound(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("not found")}}

	req := httptest.NewRequest("GET", "/translations/99", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "99"})
	w := httptest.NewRecorder()
	h.FindOne(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestFindAll_RepoError(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("db error")}}

	req := httptest.NewRequest("GET", "/translations", nil)
	w := httptest.NewRecorder()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on repo error, got none")
		}
	}()
	h.FindAll(w, req)
}

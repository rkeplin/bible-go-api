package search

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRepo struct {
	collection   TextCollection
	aggregations []SearchAggregation
	err          error
}

func (m mockRepo) Search(_, _ string, _, _ int) (TextCollection, error) {
	return m.collection, m.err
}
func (m mockRepo) SearchAggregator(_, _ string) ([]SearchAggregation, error) {
	return m.aggregations, m.err
}

func TestSearch_OK(t *testing.T) {
	collection := TextCollection{
		Total: 1,
		Items: []Text{
			{ID: 1001001, ChapterID: 1, VerseID: 1, Verse: "In the <span>beginning</span>..."},
		},
	}
	h := Handler{repo: mockRepo{collection: collection}}

	req := httptest.NewRequest("GET", "/search?query=beginning", nil)
	w := httptest.NewRecorder()
	h.Search(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got TextCollection
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if got.Total != 1 {
		t.Errorf("expected total 1, got %d", got.Total)
	}
	if len(got.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(got.Items))
	}
}

func TestSearch_Error(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("es error")}}

	req := httptest.NewRequest("GET", "/search?query=test", nil)
	w := httptest.NewRecorder()
	h.Search(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestSearch_WithOffsetAndLimit(t *testing.T) {
	h := Handler{repo: mockRepo{collection: TextCollection{Total: 0}}}

	req := httptest.NewRequest("GET", "/search?query=grace&offset=10&limit=25", nil)
	w := httptest.NewRecorder()
	h.Search(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestSearchAggregator_OK(t *testing.T) {
	aggregations := []SearchAggregation{
		{Book: Book{ID: 1, Name: "Genesis", Testament: "OT"}, Hits: 5},
		{Book: Book{ID: 43, Name: "John", Testament: "NT"}, Hits: 12},
	}
	h := Handler{repo: mockRepo{aggregations: aggregations}}

	req := httptest.NewRequest("GET", "/searchAggregator?query=love", nil)
	w := httptest.NewRecorder()
	h.SearchAggregator(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got []SearchAggregation
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 aggregations, got %d", len(got))
	}
}

func TestSearchAggregator_Error(t *testing.T) {
	h := Handler{repo: mockRepo{err: errors.New("es error")}}

	req := httptest.NewRequest("GET", "/searchAggregator?query=test", nil)
	w := httptest.NewRecorder()
	h.SearchAggregator(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

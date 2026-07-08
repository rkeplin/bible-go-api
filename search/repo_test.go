package search

import (
	"strings"
	"testing"
)

func TestGetESSearchQuery_ContainsQuery(t *testing.T) {
	r := Repository{}
	q := r.getESSearchQuery("grace", 0, 10)
	if !strings.Contains(q, "grace") {
		t.Errorf("expected query to contain 'grace', got: %s", q)
	}
	if !strings.Contains(q, `"from":0`) {
		t.Errorf("expected 'from':0 in query, got: %s", q)
	}
	if !strings.Contains(q, `"size":10`) {
		t.Errorf("expected 'size':10 in query, got: %s", q)
	}
}

func TestGetESSearchQuery_DefaultLimit(t *testing.T) {
	r := Repository{}
	q := r.getESSearchQuery("love", 0, 0)
	if !strings.Contains(q, `"size":100`) {
		t.Errorf("expected default size 100, got: %s", q)
	}
}

func TestGetESSearchQuery_MaxLimit(t *testing.T) {
	r := Repository{}
	q := r.getESSearchQuery("faith", 0, 9999)
	if !strings.Contains(q, `"size":1000`) {
		t.Errorf("expected capped size 1000, got: %s", q)
	}
}

func TestGetESSearchQuery_NegativeOffset(t *testing.T) {
	r := Repository{}
	q := r.getESSearchQuery("peace", -5, 10)
	if !strings.Contains(q, `"from":0`) {
		t.Errorf("expected offset clamped to 0, got: %s", q)
	}
}

func TestGetESSearchQuery_ContainsHighlight(t *testing.T) {
	r := Repository{}
	q := r.getESSearchQuery("hope", 0, 10)
	if !strings.Contains(q, "highlight") {
		t.Errorf("expected highlight in query, got: %s", q)
	}
}

func TestGetESSearchAggregationQuery_ContainsQuery(t *testing.T) {
	r := Repository{}
	q := r.getESSearchAggregationQuery("mercy")
	if !strings.Contains(q, "mercy") {
		t.Errorf("expected query to contain 'mercy', got: %s", q)
	}
	if !strings.Contains(q, "hitsPerBook") {
		t.Errorf("expected 'hitsPerBook' aggregation, got: %s", q)
	}
	if !strings.Contains(q, `"size": 0`) {
		t.Errorf("expected top-level size 0, got: %s", q)
	}
}

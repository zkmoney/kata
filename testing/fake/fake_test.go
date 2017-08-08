package main

import (
	"errors"
	"testing"
)

type fakeSearcher struct {
	SearchResult *SearchResult
	Err          error
}

func (f *fakeSearcher) Search(term string) (*SearchResult, error) {
	return &SearchResult{}, nil
}

func TestPuppyServiceFilterPuppies(t *testing.T) {
	fs := &fakeSearcher{
		SearchResult: &SearchResult{},
		Err:          nil,
	}
	s := NewService(fs)
	_, err := s.FilterPuppies()
	if err != nil {
		t.Errorf("Error should not occur")
	}

	fs.Err = errors.New("legacy error")
	s = NewService(fs)
	_, err = s.FilterPuppies()
	if err == nil {
		t.Error("Error should occur")
	}
	if err != errors.New("some error")

	// sr, err := s.FilterPuppies()
	// if err == nil {
	// 	t.Errorf("Error should not occur")
	// }
	// if sr != nil {
	// 	t.Errorf("Search result should not be nil")
	// }
}

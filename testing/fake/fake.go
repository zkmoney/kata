package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SearchResult struct {
	Content []byte
}

type Searcher interface {
	Search(string) (*SearchResult, error)
}

type Google struct{}

func (*Google) Search(term string) (*SearchResult, error) {
	resp, err := http.Get(fmt.Sprintf("https://google.com/#q=%s", term))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Some parsing logic

	sr := SearchResult{Content: body}
	return &sr, nil
}

type Bing struct{}

func (*Bing) Search(term string) (*SearchResult, error) {
	return nil, errors.New("don't use bing")
}

type PuppyService struct {
	client Searcher
}

func NewService(client Searcher) *PuppyService {
	return &PuppyService{client: client}
}

func (s PuppyService) FilterPuppies() (*SearchResult, error) {
	sr, err := s.client.Search("puppies")
	if err != nil {
		return nil, err
	}

	// Do a bunch of filtering logic

	return sr, nil
}

func main() {
	// s := NewService(&Google{})
	s := NewService(&Bing{})
	sr, err := s.FilterPuppies()
	fmt.Printf("%s %s", sr, err)
}

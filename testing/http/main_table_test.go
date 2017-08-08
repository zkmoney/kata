package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYellHandler_table(t *testing.T) {
	ts := httptest.NewServer(yellHandler)
	defer ts.Close()

	tests := []struct {
		In   []byte
		Want []byte
	}{
		{
			In:   []byte("soft"),
			Want: []byte("SOFT"),
		},
		{
			In:   []byte("LOUD"),
			Want: []byte("LOUD"),
		},
		{
			In:   []byte("..."),
			Want: []byte("..."),
		},
		{
			In:   []byte("SUP yo!"),
			Want: []byte("SUP YO!"),
		},
	}
	for _, test := range tests {
		resp, err := http.Post(ts.URL, "", bytes.NewBuffer(test.In))
		if err != nil {
			t.Errorf("err should be nil")
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Errorf("Successful POST should return status code 200, instead got %d", resp.StatusCode)
		}
		checkBody(t, resp.Body, test.Want)
	}
}

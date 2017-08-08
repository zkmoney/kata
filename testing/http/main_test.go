package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"io/ioutil"
)

func TestYellHandler(t *testing.T) {
	ts := httptest.NewServer(yellHandler)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("err should be nil")
	}
	if resp.StatusCode != 400 {
		t.Errorf("GET should return status code 400, instead got %d", resp.StatusCode)
	}

	in := []byte("softly spoken")
	want := []byte("SOFTLY SPOKEN")
	resp, err = http.Post(ts.URL, "", bytes.NewBuffer(in))
	if err != nil {
		t.Errorf("err should be nil")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Successful POST should return status code 200, instead got %d", resp.StatusCode)
	}
	checkBody(t, resp.Body, want)
}

func checkBody(t testing.TB, body io.Reader, want []byte) {
	have, _ := ioutil.ReadAll(body)
	if !bytes.Equal(want, have) {
		t.Errorf("Have body %s, want body %s", have, want)
	}
}

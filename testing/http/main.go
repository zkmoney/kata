package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("yell", yellHandler)

	http.ListenAndServe(port, nil)
}

var yellHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(bytes.ToUpper(body))
}

package main

import (
	"io"
	"net/http"
)
type HTTPServer struct {
	port string
}

func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

func (s HTTPServer) Open() error {
	http.HandleFunc("/", home)
	http.ListenAndServe(s.port, nil)

	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}
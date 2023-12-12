package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	s := &server{}
	http.HandleFunc("/", s.handleIndex)
	http.ListenAndServe(":8080", nil)
}

type server struct {
	t t
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	now := s.t.Now()
	fmt.Fprintf(w, "now: %v", now)
}

type t struct{}

func (t *t) Now() time.Time {
	return time.Now()
}

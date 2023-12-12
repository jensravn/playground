package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := &server{}
	http.HandleFunc("/", s.adminOnly(s.handleIndex))
	http.ListenAndServe(":8080", nil)
}

type server struct {
}

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello admin")
}

func isAdmin(r *http.Request) bool {
	return true
}

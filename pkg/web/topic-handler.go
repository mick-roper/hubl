package web

import "net/http"

// NewTopicHandler creates a new topic handler
func NewTopicHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(501)
		w.Write([]byte("not implemented"))
	}
}

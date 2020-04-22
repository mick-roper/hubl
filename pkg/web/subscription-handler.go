package web

import "net/http"

// NewSubscriptionHandler creates a new subscription handler
func NewSubscriptionHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(501)
		w.Write([]byte("not implemented"))
	}
}

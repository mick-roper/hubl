package web

import "net/http"

// SubscriptionHandler handles subscription requests
func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("not implemented"))
}

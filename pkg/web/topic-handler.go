package web

import "net/http"

// TopicHandler handles topic requests
func TopicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("not implemented"))
}

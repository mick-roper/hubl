package web

import (
	"net/http"

	"github.com/mick-roper/hubl/pkg/common"
)

// NewSubscriptionHandler creates a new subscription handler
func NewSubscriptionHandler(store common.SubscriptionStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		default:
			{
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("not implemented"))
			}
		}
	}
}

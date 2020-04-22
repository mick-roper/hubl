package web

import (
	"encoding/json"
	"net/http"

	"github.com/mick-roper/hubl/pkg/common"
)

// NewSubscriptionHandler creates a new subscription handler
func NewSubscriptionHandler(store common.SubscriptionStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			{
				subs := store.GetAll()
				bytes, err := json.Marshal(subs)

				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(fmt.SPrintf("{ \"error\": \"%v\" }", err.Error())))
					return
				}

				w.WriteHeader(200)
				w.Write(bytes)
			}
		default:
			{
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte("not implemented"))
			}
		}
	}
}

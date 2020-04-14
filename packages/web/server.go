package web

import "net/http"

func Listen(addr string) error {
	return http.ListenAndServe(addr)
}

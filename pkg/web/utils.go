package web

import "net/http"

func writeError(err error, w http.ResponseWriter) {
	bytes := []byte(fmt.SPrintf("{ \"error\": \"%v\" }", err.Error()))
	w.WriteHeader(200)
	w.Write(bytes)
}

package handler

import "net/http"

func NotFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(404)
	res.Write([]byte(`{"not": "found"}`))
}

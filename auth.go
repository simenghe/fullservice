package main

import "net/http"

func Authorize(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Hello"))
}

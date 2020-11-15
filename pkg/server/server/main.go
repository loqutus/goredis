package server

import "net/http"

func Serve() {
	http.HandleFunc("/", root)
}

package server

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hui\n")
}

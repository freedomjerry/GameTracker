package main

import (
	"fmt"
	"net/http"
)

//func ListenAndServe(addr string, handle cgi.Handler) error
//
//type Handler interface {
//	ServeHttp(ResponseWriter, *Request)
//}
func PlayerServer(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return
	}
	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return
	}

}
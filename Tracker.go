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
type PlayerStore interface {
	GetPlayerScore(name string) int
}
type PlayerServer struct {
	store PlayerStore
}

type StuPlayerStore struct {
	score map[string]int
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (s *StuPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}
//
//func GetPlayerScore(name string) (score string) {
//	if name == "Pepper"{
//		return "20"
//	}
//	if name == "Floyd"{
//		return "10"
//	}
//	return ""
//}
package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
}

type StubPlayerStore struct {
	score map[string]int
	winCalls []string
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, name string)  {
	score := p.store.GetPlayerScore(name)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
func (p *PlayerServer) processWin(w http.ResponseWriter, name string)  {
	p.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.score[name]
	return score
}
func (s *StubPlayerStore) RecordWin(name string)  {
	s.winCalls = append(s.winCalls, name)
}

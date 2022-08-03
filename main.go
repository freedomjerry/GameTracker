package main

import (
	"log"
	"net/http"
)
type InMeroryPlayerStore struct {}

func (i *InMeroryPlayerStore) GetPlayerScore(name string) int  {
	return 123
}
func (i *InMeroryPlayerStore) RecordWin (name string)  {

}
func main()  {
	server := &PlayerServer{&InMeroryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf(" could not listen on port 5000 %v ", err)
	}

}

package main


func NewInMemoryPlayerStore() *InMeroryPlayerStore {
	return &InMeroryPlayerStore{map[string]int{}}

}
type InMeroryPlayerStore struct {
	store map[string]int
}

func (i *InMeroryPlayerStore) GetPlayerScore(name string) int  {
	return i.store[name]
}
func (i *InMeroryPlayerStore) RecordWin (name string)  {
	i.store[name]++
}
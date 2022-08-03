package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoreWins(t *testing.T)  {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1{
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})

	t.Run("it record wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 2 {
			t.Fatalf("got %d, calss to RecordWin want %d", len(store.winCalls), 2)
		}
		if store.winCalls[1] != player {
			t.Errorf("did not strore correct winner got %s, want %s", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req

}
package poker_test

import (
	poker "github.com/GameTracker"
	"strings"
	"testing"
)


//func TestCLI(t *testing.T)  {
//	playerStore := &StubPlayerStore{}
//	cli := &CLI{playerStore}
//	cli.PlayPoker()
//
//	if len(playerStore.winCalls) != 1 {
//		t.Fatalf("expected a win call but didn't get any")
//	}
//
//}

func TestCLI2(t *testing.T)  {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")

	})
}

//func AssertPlayerWin(t * testing.T, store *StubPlayerStore, winner string)  {
//	t.Helper()
//
//	if len(store.winCalls) != 1 {
//		t.Fatal("expected a win call but didn't get any")
//	}
//
//	if store.winCalls[0] != winner {
//		t.Errorf("did not store correct winner got %s, want %s", store.winCalls[0], winner)
//	}
//
//}
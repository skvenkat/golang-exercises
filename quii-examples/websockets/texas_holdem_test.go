package poker_test

import (

	"fmt"	"io"

	"testing"

	"time"
)

func TestGame_Finish(t *testing.T) {

	store := &poker.StubPlayerStore{}	game := poker.NewTexasHoldem(dummyBlindAlerter, store)

	winner := "Ruth"

	game.Finish(winner)

	poker.AssertPlayerWin(t, store, winner)

}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {

	for i, want := range cases {

		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= i {

				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)

			}

			got := blindAlerter.Alerts[i]

			assertScheduledAlert(t, got, want)

		})

	}

}

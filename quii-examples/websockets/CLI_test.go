package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	 poker "github.com/skvenkat/golan-exercises/quii-examples/websockets"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled			bool
	StartCalledWith		int
	BlindAlert			[]bytes
	FinishedCalled		bool
	FinishCalledWith	string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishCalledWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {
	t.Run("start with 3 players and finish with 'Chris' as winner", func(t *testing.T){
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		
		poker.NewCLI(in, out, game)/PlayerPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T){
		game := &GameSpy{}
		in := userSends("8", "Cleo wins")
		poker.NewCLI(in, dummyStdOut, game).PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})
}


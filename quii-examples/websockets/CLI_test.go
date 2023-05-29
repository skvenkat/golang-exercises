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
	
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("pies")

		poker.NewCLI(in, out, game).PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("8", "Lloyd is a killer")

		poker.NewCLI(in, out, game).PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})	
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartCalledWith == numberOfPlayersWanted
	})

	if !passed {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

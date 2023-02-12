package main

import "fmt"

type team struct {
	A1Score int
	A2Score int
	A3Score int
	B1Score int
	B2Score int
	B3Score int
}

type ScoreCard interface {
	getScore()
}

func (t team) getScore() string {
	return fmt.Sprintf("%d%d%d:%d%d%d", t.A1Score, t.A2Score, t.A3Score, t.B1Score, t.B2Score, t.B3Score)

}

func main() {

	tournamentTeams := team{0, 0, 0, 0, 0, 0}

	scoresList := [][]int{
		{2, 3, 4, 1, 0, 3},
		{8, 4, 1, 2, 5, 7},
		{2, 0, 9, 3, 4, 4},
	}
	//for _, v in scoresList:

	tournamentTeams.getScore()

}

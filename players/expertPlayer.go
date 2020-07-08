package players

import (
    "math/rand"
    "time"
)

type expertPlayer struct {
    Strategy []int
    Wins int
    Losses int
}

func newExpertPlayer() (* expertPlayer) {
    rand.Seed(time.Now().UnixNano())
    return &expertPlayer{
        Strategy : []int{-1, -1, 5, 5, 5, -1, 9, 9, 9, -1, 13, 13, 13, -1, 17, 17, 17, -1, 21, 21, 21},
    }
}

func (p * expertPlayer) getMove(pos int) (int) {

    if p.Strategy[pos] == -1 {
        return pos + rand.Intn(3) + 1
    } else {
        return p.Strategy[pos]
    }
}

func (p * expertPlayer) getWins() (int) {
    return p.Wins
}

func (p * expertPlayer) getLosses() (int) {
    return p.Losses
}

func (p * expertPlayer) incWins() {
    p.Wins += 1
}

func (p * expertPlayer) incLosses() {
    p.Losses += 1
}

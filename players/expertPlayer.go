package players

import (
    "math/rand"
    "time"
)

type expertPlayer struct {
    strategy []int
    wins int
}

func newExpertPlayer() (* expertPlayer) {
    rand.Seed(time.Now().UnixNano())
    return &expertPlayer{
        strategy : []int{-1, -1, 5, 5, 5, -1, 9, 9, 9, -1, 13, 13, 13, -1, 17, 17, 17, -1, 21, 21, 21},
    }
}

func (p * expertPlayer) getMove(pos int) (int) {

    if p.strategy[pos] == -1 {
        return pos + rand.Intn(3) + 1
    } else {
        return p.strategy[pos]
    }
}

func (p * expertPlayer) getWins() (int) {
    return p.wins
}

func (p * expertPlayer) incWins() {
    p.wins += 1
}

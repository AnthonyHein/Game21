package players

import (
    "math/rand"
    "time"
)

type randomPlayer struct {
    strategy [][3]float32
    wins int
}

func newRandomPlayer() (* randomPlayer) {
    rand.Seed(time.Now().UnixNano())
    return &randomPlayer{
        strategy : [][3]float32{
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
            [3]float32{1/3.0, 1/3.0, 1/3.0},
        },
    }
}

func (p * randomPlayer) getMove(pos int) (int) {
    r := rand.Float32()
    var cumsum float32
    for i, v := range p.strategy[pos] {
        cumsum += v
        if r < cumsum {
            return pos + i + 1
        }
    }
    return -1
}

func (p * randomPlayer) getWins() (int) {
    return p.wins
}

func (p * randomPlayer) incWins() {
    p.wins += 1
}

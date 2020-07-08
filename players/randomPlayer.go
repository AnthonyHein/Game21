package players

import (
    "math/rand"
    "time"
)

type randomPlayer struct {
    Strategy [][3]float32
    Wins int
}

func newRandomPlayer() (* randomPlayer) {
    rand.Seed(time.Now().UnixNano())
    return &randomPlayer{
        Strategy : [][3]float32{
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
    for i, v := range p.Strategy[pos] {
        cumsum += v
        if r < cumsum {
            return pos + i + 1
        }
    }
    return -1
}

func (p * randomPlayer) getWins() (int) {
    return p.Wins
}

func (p * randomPlayer) incWins() {
    p.Wins += 1
}

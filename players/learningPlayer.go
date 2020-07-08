package players

import (
    "encoding/json"
    "math/rand"
    "io/ioutil"
    "time"
    "log"
)

type learningPlayer struct {
    Strategy [][3]float32
    Wins int
}

func newLearningPlayer() (* learningPlayer) {
    rand.Seed(time.Now().UnixNano())

    var lp learningPlayer

    // Attempt to load data.
    bs, err := ioutil.ReadFile("learningPlayer.json")
    if err == nil {
        err = json.Unmarshal(bs, &lp)
        if err == nil {
            lp.Wins = 0
            return &lp
        }
    }

    // Did not load data.
    lp = learningPlayer{
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
    return &lp
}

func (p * learningPlayer) getMove(pos int) (int) {
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

func (p * learningPlayer) getWins() (int) {
    return p.Wins
}

func (p * learningPlayer) incWins() {
    p.Wins += 1
}

func (p * learningPlayer) save() {
    bs, err := json.MarshalIndent(p, "", "\t")
    if err != nil {
        log.Fatalln(err)
    }
    err = ioutil.WriteFile("players/learningPlayer.json", bs, 0666)
    if err != nil {
        log.Fatalln(err)
    }
}

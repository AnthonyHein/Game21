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
    ActionSequence map[int]int
    Wins int
    Losses int
}

func newLearningPlayer() (* learningPlayer) {
    rand.Seed(time.Now().UnixNano())

    var lp learningPlayer

    // Attempt to load data.
    bs, err := ioutil.ReadFile("players/learningPlayer.json")
    if err == nil {
        err = json.Unmarshal(bs, &lp)
        if err == nil {
            lp.ActionSequence = make(map[int]int)
            lp.Wins = 0
            lp.Losses = 0
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
        ActionSequence : make(map[int]int),
    }
    return &lp
}

func (p * learningPlayer) getMove(pos int) (int) {
    r := rand.Float32()
    var cumsum float32
    for i, v := range p.Strategy[pos] {
        cumsum += v
        if r < cumsum {
            p.ActionSequence[pos] = i
            return pos + i + 1
        }
    }
    print("There has been a normalization error.")
    return -1
}

func (p * learningPlayer) getWins() (int) {
    return p.Wins
}

func (p * learningPlayer) getLosses() (int) {
    return p.Losses
}

func (p * learningPlayer) incWins() {
    p.Wins += 1
    p.reward()
    p.ActionSequence = make(map[int]int)
}

func (p * learningPlayer) incLosses() {
    p.Losses += 1
    p.punish()
    p.ActionSequence = make(map[int]int)
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

func (p * learningPlayer) reward() {
    for pos, next := range p.ActionSequence {

        var boost float32
        var total float32

        for i, v := range p.Strategy[pos] {
            if i != next {
                boost += v * 0.1
                p.Strategy[pos][i] = v * 0.9
                total += p.Strategy[pos][i]
            }
        }
        p.Strategy[pos][next] += boost
        total += p.Strategy[pos][next]

        // Normalize.
        for i, _ := range p.Strategy[pos] {
            p.Strategy[pos][i] /= total
        }
    }
}

func (p * learningPlayer) punish() {
    for pos, next := range p.ActionSequence {

        var setback float32
        var total float32

        setback = p.Strategy[pos][next] * 0.1
        p.Strategy[pos][next] *= 0.9

        for i, _ := range p.Strategy[pos] {
            if i != next {
                p.Strategy[pos][i] += setback / 2.0
            }
            total += p.Strategy[pos][i]
        }

        // Normalize.
        for i, _ := range p.Strategy[pos] {
            p.Strategy[pos][i] /= total
        }

    }
}

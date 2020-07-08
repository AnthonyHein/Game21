package players

import (
    "log"
)

type PlayerType string

const(
    Expert PlayerType = "Expert"
    Random PlayerType = "Random"
    Learning PlayerType = "Learning"
)

type Player interface {
    getMove(int) (int)
    getWins() (int)
    incWins()
}

func NewPlayer(s PlayerType) (Player) {
    if s == Expert {
        return newExpertPlayer()
    } else if s == Random {
        return newRandomPlayer()
    } else if s == Learning {
        return newLearningPlayer()
    } else {
        log.Fatal("This player type is not recognized by players.go: ", string(s))
        return nil
    }
}

func GetMove(p Player, pos int) (int) {
    return p.getMove(pos)
}

func GetWins(p Player) (int) {
    return p.getWins()
}

func IncWins(p Player) {
    p.incWins()
}

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
    getLosses() (int)
    incWins()
    incLosses()
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

func GetLosses(p Player) (int) {
    return p.getLosses()
}

func IncWins(p Player) {
    p.incWins()
}

func IncLosses(p Player) {
    p.incLosses()
}

func Save(p Player) {
    if lp, ok := p.(* learningPlayer); ok {
        lp.save()
    }
}

package game

import (
    "Game21/players"
    "math/rand"
    "fmt"
    "strconv"
)

type Game struct {
    ps [2]players.Player
}

func New(p1, p2 players.Player) (* Game) {
    return &Game{
        ps : [2]players.Player{p1, p2},
    }
}

func (g * Game) Simulate(n int) {
    for i := 0; i < n; i++ {

        currentPlayer := rand.Intn(2)

        var pos int
        for 0 <= pos && pos < 21 {

            pos = players.GetMove(g.ps[currentPlayer], pos)
            currentPlayer = (currentPlayer + 1) % 2

        }

        if pos == 21 {
            players.IncWins(g.ps[(currentPlayer + 1) % 2])
        } else {
            players.IncWins(g.ps[currentPlayer])
        }
    }
}

func (g * Game) Results() {
    fmt.Println("Player 1 wins: " + strconv.Itoa(players.GetWins(g.ps[0])))
    fmt.Println("Player 2 wins: " + strconv.Itoa(players.GetWins(g.ps[1])))
}

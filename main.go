package main

import (
    "Game21/players"
    "Game21/game"
)

func main () {
    p1 := players.NewPlayer(players.Expert)
    p2 := players.NewPlayer(players.Learning)

    g := game.New(p1, p2)
    g.Simulate(1000000)
    g.Results()
}

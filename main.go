package main

import (
    "Game21/players"
    "Game21/game"
)

func main () {
    p1 := players.NewPlayer(players.Expert)
    p2 := players.NewPlayer(players.Random)

    g := game.New(p1, p2)
    g.Simulate(1000)
    g.Results()
}
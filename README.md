### Simple Game RL Agent

This repository implements an RL agent written in Golang which learns how to
play a simple game.

### What's the game?

This is a two player game where the objective is to be the first to reach 21.
The count starts at 0 and players take turns increasing the count by 1, 2, or 3
points at a time.

Sample gameplay:

* Anthony starts at 0 and adds 3 (3).
* Michael starts at 3 and adds 3 (6).
* Anthony starts at 6 and adds 2 (8).
* Michael starts at 8 and adds 1 (9).
* Anthony starts at 9 and adds 3 (12).
* Michael starts at 12 and adds 2 (14).
* Anthony starts at 14 and adds 1 (15).
* Michael starts at 15 and adds 3 (18).
* Anthony starts at 18 and adds 3 (21).

Anthony wins!

### Winning Strategy

Surprisingly, there is a very real and proven strategy for this simple game that,
if followed by the first player to go, will always result in that player winning
and if followed by the second player to go and unknown to the first player to go
will usually result in a win (my tests have shown that an 'ExpertPlayer' wins
999 games out of 1000 against a 'RandomPlayer').

Do you know what the strategy is?

### Results

With about ```1,000,000``` simulations, the 'LearningPlayer' goes from randomly
selecting actions to perfectly executing the optimal strategy at all points.

Cool stuff!

### Usage

1. Of course, make sure you have Go installed. Check out the offical site [here](https://golang.org/).

2. In terminal
```console
$ go build -o mygame
$ ./mygame
```

### File Structure

```
.
├── mygame (executable)
├── README.md (you are reading this right now)
├── game
│   └── game.go (dictates how the game is played)
├── main.go (creates players and runs games)
└── players
    ├── player.go (implements Player interface and functions)
    ├── learningPlayer.go (the RL agent which learns the game as it goes along)
    ├── expertPlayer.go (the agent which knows the winning strategy and plays accordingly)
    └── randomPlayer.go (the agent which takes random moves)
```

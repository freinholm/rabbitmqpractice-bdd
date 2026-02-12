package main

import (
	"fmt"

	"github.com/freinholm/rabbitmqpractice-bdd/internal/gamelogic"
	"github.com/freinholm/rabbitmqpractice-bdd/internal/routing"
)

func handlerMove(gs *gamelogic.GameState) func(gamelogic.ArmyMove) {
	return func(move gamelogic.ArmyMove) {
		defer fmt.Print("> ")
		gs.HandleMove(move)
	}
}

func handlerPause(gs *gamelogic.GameState) func(routing.PlayingState) {
	return func(ps routing.PlayingState) {
		defer fmt.Print("> ")
		gs.HandlePause(ps)
	}
}

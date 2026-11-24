package main

import (
	"log"

	game "github.com/andrewbakradze/life-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	board, err := game.NewBoard(100, 100)
	if err != nil {
		log.Printf("board initializing error: %s", err.Error())
		return
	}
	game := &game.LifeGame{Board: board}
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Simple life game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	game "github.com/andrewbakradze/life-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

var deffoultBoardSize = 20

func main() {

	board, err := game.NewBoard(deffoultBoardSize, deffoultBoardSize)
	if err != nil {
		log.Printf("board initializing error: %s", err.Error())
		return
	}
	game := game.NewGame(board)
	ebiten.SetTPS(30)
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Simple life game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	backgroundColor = color.RGBA{R: 94, G: 2, B: 2, A: 255}
	aliveColor      = color.RGBA{R: 206, G: 2, B: 2, A: 255}
)

type LifeGame struct {
	started bool
	Board   *board
}

func NewGame(board *board) *LifeGame {
	return &LifeGame{Board: board}
}

func (g *LifeGame) invalidCords(x, y int) bool {
	return x < 0 || y < 0 || x >= g.Board.width || y >= g.Board.height
}

func (g *LifeGame) changeMode() {
	g.started = !g.started
}

func (g *LifeGame) Update() error {
	if !g.started && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.updateCell(g.Board.SetCell)
	} else if !g.started && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.updateCell(g.Board.ResetCell)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.changeMode()
	}
	if g.started {
		g.updateBoard()
	}
	return nil
}

func (g *LifeGame) updateCell(updateFunction func(int, int)) {
	x, y := ebiten.CursorPosition()
	if g.invalidCords(x, y) {
		return
	}
	updateFunction(x, y)
	log.Printf("changed cell x = %v, y = %v\n", x, y)
}

func (g *LifeGame) countNeiboors(x, y int) int {
	neighbors := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if g.invalidCords(nx, ny) {
				continue
			}
			if g.Board.GetCell(nx, ny) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (g *LifeGame) updateBoard() {
	next, _ := NewBoard(g.Board.width, g.Board.height)
	for x := range g.Board.width {
		for y := range g.Board.height {
			neighbors := g.countNeiboors(x, y)
			if !(neighbors == 3 || (g.Board.GetCell(x, y) && neighbors == 2)) {
				continue
			}
			next.SetCell(x, y)
		}
	}
	g.Board = next
}

func (g *LifeGame) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	for i, alive := range g.Board.cells {
		if alive {
			screen.Set(i%g.Board.width, i/g.Board.width, aliveColor)
		}
	}
}

func (g *LifeGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Board.height, g.Board.width
}

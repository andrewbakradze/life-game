package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	backgroundColor = color.RGBA{R: 94, G: 2, B: 2, A: 1}
	aliveColor      = color.RGBA{R: 206, G: 2, B: 2, A: 1}
	boardSize       = 20
)

type LifeGame struct {
	started bool
	Board   *board
}

func invalidCord(cord int) bool {
	return cord < 0 || cord >= boardSize
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
	if invalidCord(x) || invalidCord(y) {
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
			if nx < 0 || nx >= g.Board.width || ny < 0 || ny >= g.Board.height {
				continue
			}
			if g.Board.cells[ny*g.Board.width+nx] {
				neighbors++
			}
		}
	}
	return neighbors
}

func (g *LifeGame) updateBoard() {

	next := make([]bool, len(g.Board.cells))

	for i, alive := range g.Board.cells {
		x, y := i%g.Board.width, i/g.Board.width
		neighbors := g.countNeiboors(x, y)
		next[i] = neighbors == 3 || (alive && neighbors == 2)
	}

	g.Board.cells = next
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
	return boardSize, boardSize
}

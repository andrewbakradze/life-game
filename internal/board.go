package game

import "errors"

type board struct {
	width  int
	height int
	cells  []bool
}

var (
	ErrInvalidBoardSize = errors.New("board width and height must be positive")
)

func NewBoard(width, height int) (*board, error) {
	if width <= 0 || height <= 0 {
		return nil, ErrInvalidBoardSize
	}
	return &board{width: width, height: height, cells: make([]bool, width*height)}, nil
}

func (b *board) ChangeCell(x, y int) {
	b.cells[y*b.width+x] = !b.cells[y*b.width+x]
}

func (b *board) GetCell(x, y int) bool {
	return b.cells[y*b.width+x]
}

func (b *board) SetCell(x, y int) {
	b.cells[y*b.width+x] = true
}

func (b *board) ResetCell(x, y int) {
	b.cells[y*b.width+x] = false
}

package board

import "github.com/thunderlight-shogi/engine/internal/model"

type board struct {
	Cells [][]*Figure // [Horizontal offset][Vertical offset]
}

type Board = *board

type Figure struct {
	Type             model.FigureType
	HorizontalOffset uint
	VerticalOffset   uint
}

func (this_board Board) Clone() (newby Board) {
	newby = new(board)
	newby.Cells = make([][]*Figure, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Figure, 9)
		copy(newby.Cells[i], this_board.Cells[i])
	}
	return
}

func Construct() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Figure, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Figure, 9)
	}
	return
}

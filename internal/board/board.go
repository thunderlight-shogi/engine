package board

import "github.com/thunderlight-shogi/engine/internal/model"

type Board struct {
	Cells [][]*Figure // [Horizontal offset][Vertical offset]
}

type Figure struct {
	Type             model.FigureType
	HorizontalOffset uint
	VerticalOffset   uint
}

func (board *Board) Clone() (newby *Board) {
	newby = new(Board)
	newby.Cells = make([][]*Figure, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Figure, 9)
	}
	return
}

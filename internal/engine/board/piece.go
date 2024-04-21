package board

import "github.com/thunderlight-shogi/engine/internal/model"

type Piece struct {
	Type   *model.PieceType
	Player model.Player
}

func (this_piece *Piece) IsPromoted() bool {
	return this_piece.Type.DemotePiece != nil
}

func (this_piece *Piece) IsPromotable() bool {
	return this_piece.Type.PromotePiece != nil
}

func (this_piece *Piece) getShiftSign() int {
	if this_piece.Player == model.Sente {
		return 1
	} else {
		return -1
	}
}

func (this_piece *Piece) GetPromotedPiece() *Piece {
	return &Piece{Type: this_piece.Type.PromotePiece, Player: this_piece.Player}
}

func (this_piece *Piece) GetAttackerPlayer() (attackerPlayer model.Player) {
	if this_piece.Player == model.Sente {
		attackerPlayer = model.Gote
	} else {
		attackerPlayer = model.Sente
	}
	return
}

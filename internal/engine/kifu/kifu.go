package kifu

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
)

var ErrGameAlreadyFinished = errors.New("game already finished")

type KifuWriter struct {
	prev_pos      board.Position
	buffer        io.Writer
	move_number   uint
	game_finished bool
}

type InitStruct struct {
	SenteName string
	GoteName  string
}

var kanji_numerals = [...]string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}

func Create(writer io.Writer) (new_writer KifuWriter) {
	new_writer.prev_pos = board.NewPos(-1, -1)
	new_writer.buffer = writer
	new_writer.move_number = 0
	new_writer.game_finished = false
	return
}

// ：

func (writer KifuWriter) Init(str InitStruct) error {
	_, err := fmt.Fprintf(writer.buffer, "先手:%s\n後手:%s\n手数----指手---------消費時間--\n",
		str.SenteName, str.GoteName)
	return err
}

func (writer *KifuWriter) Move(move board.Move, elapsed time.Duration) error {
	writer.move_number++
	var (
		from    string
		to      string
		piece   string
		mv_type string
		mv      string
	)

	if writer.game_finished {
		return ErrGameAlreadyFinished
	}

	if move.MoveType == board.Surrender {
		mv = "投了"
		writer.game_finished = true
	} else {

		if move.NewCoords == writer.prev_pos {
			to = "同 "
		} else {
			to = fmt.Sprintf("%d%s", move.NewCoords.File+1, kanji_numerals[move.NewCoords.Rank])
		}
		piece = string(move.PieceType.Kanji)

		switch move.MoveType {
		case board.Attacking, board.Moving:
			mv_type = ""
			from = fmt.Sprintf("(%d%d)", move.OldCoords.File+1, move.OldCoords.Rank+1)
		case board.PromotionAttacking, board.PromotionMoving:
			mv_type = "成"
			from = fmt.Sprintf("(%d%d)", move.OldCoords.File+1, move.OldCoords.Rank+1)
		case board.Dropping:
			mv_type = "打"
			from = ""
		}

		mv = to + piece + mv_type + from
	}

	writer.prev_pos = move.NewCoords

	var (
		minutes uint = uint(elapsed.Minutes())
		seconds uint = uint(elapsed.Seconds()) - 60*minutes
	)

	_, err := fmt.Fprintf(writer.buffer, "%4d   %-8s (%02d:%02d)\n",
		writer.move_number,
		mv,
		minutes,
		seconds,
	)
	return err
}

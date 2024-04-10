package model

func seed() {
	// Silver+
	var SilverPlus = PieceType{
		Id:    12,
		Name:  "Silver+",
		Kanji: '全',
		Cost:  6,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
		},
	}

	// Knight+
	var KnightPlus = PieceType{
		Id:    10,
		Name:  "Knight+",
		Kanji: '圭',
		Cost:  6,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
		},
	}

	// Lance+
	var LancePlus = PieceType{
		Id:    8,
		Name:  "Lance+",
		Kanji: '杏',
		Cost:  6,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
		},
	}

	// Pawn+
	var PawnPlus = PieceType{
		Id:    2,
		Name:  "Pawn+",
		Kanji: 'と',
		Cost:  7,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
		},
	}
	// Rook+
	var RookPlus = PieceType{
		Id:    6,
		Name:  "Rook+",
		Kanji: '竜',
		Cost:  11,
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -8,
			},

			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   2,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   3,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   4,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   5,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   6,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   7,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   8,
			},

			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   0,
			},

			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   0,
			},

			{
				HorizontalShift: 1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
		},
	}

	// Bishop+
	var BishopPlus = PieceType{
		Id:    4,
		Name:  "Bishop+",
		Kanji: '馬',
		Cost:  10,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   2,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   3,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   4,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   5,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   6,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   7,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   8,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   2,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   3,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   4,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   5,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   6,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   7,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   8,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   -8,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   -8,
			},

			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
		},
	}

	// King
	var King = PieceType{
		Id:             14,
		Name:           "King",
		Kanji:          '王',
		ImportantPiece: true,
		Cost:           999999,
		Moves: []Move{
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   1,
			},
		},
	}

	// Gold
	var Gold = PieceType{
		Id:    13,
		Name:  "Gold",
		Kanji: '金',
		Cost:  6,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
		},
	}

	// Silver
	var Silver = PieceType{
		Id:           11,
		Name:         "Silver",
		Kanji:        '銀',
		Cost:         5,
		PromotePiece: &SilverPlus,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   1,
			},
		},
	}

	// Knight
	var Knight = PieceType{
		Id:           9,
		Name:         "Knight",
		Kanji:        '桂',
		Cost:         4,
		PromotePiece: &KnightPlus,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -2,
			},
		},
	}

	// Lance
	var Lance = PieceType{
		Id:           7,
		Name:         "Lance",
		Kanji:        '香',
		Cost:         3,
		PromotePiece: &LancePlus,
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -8,
			},
		},
	}

	// Rook
	var Rook = PieceType{
		Id:           5,
		Name:         "Rook",
		Kanji:        '飛',
		Cost:         9,
		PromotePiece: &RookPlus,
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   -8,
			},

			{
				HorizontalShift: 0,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   2,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   3,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   4,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   5,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   6,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   7,
			},
			{
				HorizontalShift: 0,
				VerticalShift:   8,
			},

			{
				HorizontalShift: -1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   0,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   0,
			},

			{
				HorizontalShift: 1,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   0,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   0,
			},
		},
	}

	// Bishop
	var Bishop = PieceType{
		Id:           3,
		Name:         "Bishop",
		Kanji:        '角',
		Cost:         8,
		PromotePiece: &BishopPlus,
		Moves: []Move{
			{
				HorizontalShift: 1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   2,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   3,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   4,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   5,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   6,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   7,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   8,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   1,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   2,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   3,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   4,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   5,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   6,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   7,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   8,
			},
			{
				HorizontalShift: 1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: 2,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: 3,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: 4,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: 5,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: 6,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: 7,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: 8,
				VerticalShift:   -8,
			},
			{
				HorizontalShift: -1,
				VerticalShift:   -1,
			},
			{
				HorizontalShift: -2,
				VerticalShift:   -2,
			},
			{
				HorizontalShift: -3,
				VerticalShift:   -3,
			},
			{
				HorizontalShift: -4,
				VerticalShift:   -4,
			},
			{
				HorizontalShift: -5,
				VerticalShift:   -5,
			},
			{
				HorizontalShift: -6,
				VerticalShift:   -6,
			},
			{
				HorizontalShift: -7,
				VerticalShift:   -7,
			},
			{
				HorizontalShift: -8,
				VerticalShift:   -8,
			},
		},
	}

	// Pawn
	var Pawn = PieceType{
		Id:           1,
		Name:         "Pawn",
		Kanji:        '歩',
		Cost:         1,
		PromotePiece: &PawnPlus,
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
		},
	}

	var DefaultPosition = StartingPosition{
		Id:   1,
		Name: "Default",
		Pieces: []StartingPositionPiece{
			{
				PieceType:        &Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   3,
				Player:           Gote,
			},

			{
				PieceType:        &Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        &Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   7,
				Player:           Sente,
			},

			{
				PieceType:        &Rook,
				HorizontalOffset: 8,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				PieceType:        &Rook,
				HorizontalOffset: 2,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				PieceType:        &Bishop,
				HorizontalOffset: 2,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				PieceType:        &Bishop,
				HorizontalOffset: 8,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				PieceType:        &Lance,
				HorizontalOffset: 1,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Knight,
				HorizontalOffset: 2,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Silver,
				HorizontalOffset: 3,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Gold,
				HorizontalOffset: 4,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &King,
				HorizontalOffset: 5,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Gold,
				HorizontalOffset: 6,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Silver,
				HorizontalOffset: 7,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Knight,
				HorizontalOffset: 8,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        &Lance,
				HorizontalOffset: 9,
				VerticalOffset:   1,
				Player:           Gote,
			},

			{
				PieceType:        &Lance,
				HorizontalOffset: 1,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Knight,
				HorizontalOffset: 2,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Silver,
				HorizontalOffset: 3,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Gold,
				HorizontalOffset: 4,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &King,
				HorizontalOffset: 5,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Gold,
				HorizontalOffset: 6,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Silver,
				HorizontalOffset: 7,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Knight,
				HorizontalOffset: 8,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        &Lance,
				HorizontalOffset: 9,
				VerticalOffset:   9,
				Player:           Sente,
			},
		},
	}

	db.Create(&DefaultPosition)
}

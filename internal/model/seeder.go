package model

func seed() {
	// Silver+
	var SilverPlus = PieceType{
		Name:  "Silver+",
		Kanji: '全',
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
		Name:  "Knight+",
		Kanji: '圭',
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
		Name:  "Lance+",
		Kanji: '杏',
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
		Name:  "Pawn+",
		Kanji: 'と',
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
		Name:  "Rook+",
		Kanji: '竜',
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
		Name:  "Bishop+",
		Kanji: '馬',
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
		Name:           "King",
		Kanji:          '王',
		ImportantPiece: true,
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
		Name:  "Gold",
		Kanji: '金',
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
		Name:  "Silver",
		Kanji: '銀',
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
		PromotePiece: &SilverPlus,
	}

	// Knight
	var Knight = PieceType{
		Name:  "Knight",
		Kanji: '桂',
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
		PromotePiece: &KnightPlus,
	}

	// Lance
	var Lance = PieceType{
		Name:  "Lance",
		Kanji: '香',
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
		PromotePiece: &LancePlus,
	}

	// Rook
	var Rook = PieceType{
		Name:  "Rook",
		Kanji: '飛',
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
		PromotePiece: &RookPlus,
	}

	// Bishop
	var Bishop = PieceType{
		Name:  "Bishop",
		Kanji: '角',
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
		PromotePiece: &BishopPlus,
	}

	// Pawn
	var Pawn = PieceType{
		Name:  "Pawn",
		Kanji: '歩',
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
		},
		PromotePiece: &PawnPlus,
	}

	db.Create([]*PieceType{&King, &Gold, &Silver, &Knight, &Lance, &Rook, &Bishop, &Pawn})

	var DefaultPosition = StartingPosition{
		Name: "Default",
		Pieces: []StartingPositionPiece{
			{
				PieceType:        Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   3,
				Player:           Gote,
			},

			{
				PieceType:        Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				PieceType:        Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   7,
				Player:           Sente,
			},

			{
				PieceType:        Rook,
				HorizontalOffset: 8,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				PieceType:        Rook,
				HorizontalOffset: 2,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				PieceType:        Bishop,
				HorizontalOffset: 2,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				PieceType:        Bishop,
				HorizontalOffset: 8,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				PieceType:        Lance,
				HorizontalOffset: 1,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Knight,
				HorizontalOffset: 2,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Silver,
				HorizontalOffset: 3,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Gold,
				HorizontalOffset: 4,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        King,
				HorizontalOffset: 5,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Gold,
				HorizontalOffset: 6,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Silver,
				HorizontalOffset: 7,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Knight,
				HorizontalOffset: 8,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				PieceType:        Lance,
				HorizontalOffset: 9,
				VerticalOffset:   1,
				Player:           Gote,
			},

			{
				PieceType:        Lance,
				HorizontalOffset: 1,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Knight,
				HorizontalOffset: 2,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Silver,
				HorizontalOffset: 3,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Gold,
				HorizontalOffset: 4,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        King,
				HorizontalOffset: 5,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Gold,
				HorizontalOffset: 6,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Silver,
				HorizontalOffset: 7,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Knight,
				HorizontalOffset: 8,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				PieceType:        Lance,
				HorizontalOffset: 9,
				VerticalOffset:   9,
				Player:           Sente,
			},
		},
	}

	db.Create(&DefaultPosition)
}

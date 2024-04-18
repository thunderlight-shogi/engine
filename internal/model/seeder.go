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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 0,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 0,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 0,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 0,
				RankShift: 1,
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
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -2,
			},
			{
				FileShift: 0,
				RankShift: -3,
			},
			{
				FileShift: 0,
				RankShift: -4,
			},
			{
				FileShift: 0,
				RankShift: -5,
			},
			{
				FileShift: 0,
				RankShift: -6,
			},
			{
				FileShift: 0,
				RankShift: -7,
			},
			{
				FileShift: 0,
				RankShift: -8,
			},

			{
				FileShift: 0,
				RankShift: 1,
			},
			{
				FileShift: 0,
				RankShift: 2,
			},
			{
				FileShift: 0,
				RankShift: 3,
			},
			{
				FileShift: 0,
				RankShift: 4,
			},
			{
				FileShift: 0,
				RankShift: 5,
			},
			{
				FileShift: 0,
				RankShift: 6,
			},
			{
				FileShift: 0,
				RankShift: 7,
			},
			{
				FileShift: 0,
				RankShift: 8,
			},

			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: -2,
				RankShift: 0,
			},
			{
				FileShift: -3,
				RankShift: 0,
			},
			{
				FileShift: -4,
				RankShift: 0,
			},
			{
				FileShift: -5,
				RankShift: 0,
			},
			{
				FileShift: -6,
				RankShift: 0,
			},
			{
				FileShift: -7,
				RankShift: 0,
			},
			{
				FileShift: -8,
				RankShift: 0,
			},

			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: 2,
				RankShift: 0,
			},
			{
				FileShift: 3,
				RankShift: 0,
			},
			{
				FileShift: 4,
				RankShift: 0,
			},
			{
				FileShift: 5,
				RankShift: 0,
			},
			{
				FileShift: 6,
				RankShift: 0,
			},
			{
				FileShift: 7,
				RankShift: 0,
			},
			{
				FileShift: 8,
				RankShift: 0,
			},

			{
				FileShift: 1,
				RankShift: 1,
			},
			{
				FileShift: -1,
				RankShift: 1,
			},
			{
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
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
				FileShift: 1,
				RankShift: 1,
			},
			{
				FileShift: 2,
				RankShift: 2,
			},
			{
				FileShift: 3,
				RankShift: 3,
			},
			{
				FileShift: 4,
				RankShift: 4,
			},
			{
				FileShift: 5,
				RankShift: 5,
			},
			{
				FileShift: 6,
				RankShift: 6,
			},
			{
				FileShift: 7,
				RankShift: 7,
			},
			{
				FileShift: 8,
				RankShift: 8,
			},
			{
				FileShift: -1,
				RankShift: 1,
			},
			{
				FileShift: -2,
				RankShift: 2,
			},
			{
				FileShift: -3,
				RankShift: 3,
			},
			{
				FileShift: -4,
				RankShift: 4,
			},
			{
				FileShift: -5,
				RankShift: 5,
			},
			{
				FileShift: -6,
				RankShift: 6,
			},
			{
				FileShift: -7,
				RankShift: 7,
			},
			{
				FileShift: -8,
				RankShift: 8,
			},
			{
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 2,
				RankShift: -2,
			},
			{
				FileShift: 3,
				RankShift: -3,
			},
			{
				FileShift: 4,
				RankShift: -4,
			},
			{
				FileShift: 5,
				RankShift: -5,
			},
			{
				FileShift: 6,
				RankShift: -6,
			},
			{
				FileShift: 7,
				RankShift: -7,
			},
			{
				FileShift: 8,
				RankShift: -8,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: -2,
				RankShift: -2,
			},
			{
				FileShift: -3,
				RankShift: -3,
			},
			{
				FileShift: -4,
				RankShift: -4,
			},
			{
				FileShift: -5,
				RankShift: -5,
			},
			{
				FileShift: -6,
				RankShift: -6,
			},
			{
				FileShift: -7,
				RankShift: -7,
			},
			{
				FileShift: -8,
				RankShift: -8,
			},

			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: 1,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
		},
	}

	// King
	var King = PieceType{
		Id:             14,
		Name:           "King",
		Kanji:          '王',
		ImportantPiece: true,
		Cost:           0,
		Moves: []Move{
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: 1,
			},
			{
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: 1,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: 0,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: 1,
				RankShift: 1,
			},
			{
				FileShift: -1,
				RankShift: 1,
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
				FileShift: 1,
				RankShift: -2,
			},
			{
				FileShift: -1,
				RankShift: -2,
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
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -2,
			},
			{
				FileShift: 0,
				RankShift: -3,
			},
			{
				FileShift: 0,
				RankShift: -4,
			},
			{
				FileShift: 0,
				RankShift: -5,
			},
			{
				FileShift: 0,
				RankShift: -6,
			},
			{
				FileShift: 0,
				RankShift: -7,
			},
			{
				FileShift: 0,
				RankShift: -8,
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
				FileShift: 0,
				RankShift: -1,
			},
			{
				FileShift: 0,
				RankShift: -2,
			},
			{
				FileShift: 0,
				RankShift: -3,
			},
			{
				FileShift: 0,
				RankShift: -4,
			},
			{
				FileShift: 0,
				RankShift: -5,
			},
			{
				FileShift: 0,
				RankShift: -6,
			},
			{
				FileShift: 0,
				RankShift: -7,
			},
			{
				FileShift: 0,
				RankShift: -8,
			},

			{
				FileShift: 0,
				RankShift: 1,
			},
			{
				FileShift: 0,
				RankShift: 2,
			},
			{
				FileShift: 0,
				RankShift: 3,
			},
			{
				FileShift: 0,
				RankShift: 4,
			},
			{
				FileShift: 0,
				RankShift: 5,
			},
			{
				FileShift: 0,
				RankShift: 6,
			},
			{
				FileShift: 0,
				RankShift: 7,
			},
			{
				FileShift: 0,
				RankShift: 8,
			},

			{
				FileShift: -1,
				RankShift: 0,
			},
			{
				FileShift: -2,
				RankShift: 0,
			},
			{
				FileShift: -3,
				RankShift: 0,
			},
			{
				FileShift: -4,
				RankShift: 0,
			},
			{
				FileShift: -5,
				RankShift: 0,
			},
			{
				FileShift: -6,
				RankShift: 0,
			},
			{
				FileShift: -7,
				RankShift: 0,
			},
			{
				FileShift: -8,
				RankShift: 0,
			},

			{
				FileShift: 1,
				RankShift: 0,
			},
			{
				FileShift: 2,
				RankShift: 0,
			},
			{
				FileShift: 3,
				RankShift: 0,
			},
			{
				FileShift: 4,
				RankShift: 0,
			},
			{
				FileShift: 5,
				RankShift: 0,
			},
			{
				FileShift: 6,
				RankShift: 0,
			},
			{
				FileShift: 7,
				RankShift: 0,
			},
			{
				FileShift: 8,
				RankShift: 0,
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
				FileShift: 1,
				RankShift: 1,
			},
			{
				FileShift: 2,
				RankShift: 2,
			},
			{
				FileShift: 3,
				RankShift: 3,
			},
			{
				FileShift: 4,
				RankShift: 4,
			},
			{
				FileShift: 5,
				RankShift: 5,
			},
			{
				FileShift: 6,
				RankShift: 6,
			},
			{
				FileShift: 7,
				RankShift: 7,
			},
			{
				FileShift: 8,
				RankShift: 8,
			},
			{
				FileShift: -1,
				RankShift: 1,
			},
			{
				FileShift: -2,
				RankShift: 2,
			},
			{
				FileShift: -3,
				RankShift: 3,
			},
			{
				FileShift: -4,
				RankShift: 4,
			},
			{
				FileShift: -5,
				RankShift: 5,
			},
			{
				FileShift: -6,
				RankShift: 6,
			},
			{
				FileShift: -7,
				RankShift: 7,
			},
			{
				FileShift: -8,
				RankShift: 8,
			},
			{
				FileShift: 1,
				RankShift: -1,
			},
			{
				FileShift: 2,
				RankShift: -2,
			},
			{
				FileShift: 3,
				RankShift: -3,
			},
			{
				FileShift: 4,
				RankShift: -4,
			},
			{
				FileShift: 5,
				RankShift: -5,
			},
			{
				FileShift: 6,
				RankShift: -6,
			},
			{
				FileShift: 7,
				RankShift: -7,
			},
			{
				FileShift: 8,
				RankShift: -8,
			},
			{
				FileShift: -1,
				RankShift: -1,
			},
			{
				FileShift: -2,
				RankShift: -2,
			},
			{
				FileShift: -3,
				RankShift: -3,
			},
			{
				FileShift: -4,
				RankShift: -4,
			},
			{
				FileShift: -5,
				RankShift: -5,
			},
			{
				FileShift: -6,
				RankShift: -6,
			},
			{
				FileShift: -7,
				RankShift: -7,
			},
			{
				FileShift: -8,
				RankShift: -8,
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
				FileShift: 0,
				RankShift: -1,
			},
		},
	}

	var DefaultPreset = Preset{
		Id:   1,
		Name: "Default",
		Pieces: []PresetPiece{
			{
				PieceType: &Pawn,
				File:      1,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      2,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      3,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      4,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      5,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      6,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      7,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      8,
				Rank:      3,
				Player:    Gote,
			},
			{
				PieceType: &Pawn,
				File:      9,
				Rank:      3,
				Player:    Gote,
			},

			{
				PieceType: &Pawn,
				File:      1,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      2,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      3,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      4,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      5,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      6,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      7,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      8,
				Rank:      7,
				Player:    Sente,
			},
			{
				PieceType: &Pawn,
				File:      9,
				Rank:      7,
				Player:    Sente,
			},

			{
				PieceType: &Rook,
				File:      8,
				Rank:      2,
				Player:    Gote,
			},
			{
				PieceType: &Rook,
				File:      2,
				Rank:      8,
				Player:    Sente,
			},

			{
				PieceType: &Bishop,
				File:      2,
				Rank:      2,
				Player:    Gote,
			},
			{
				PieceType: &Bishop,
				File:      8,
				Rank:      8,
				Player:    Sente,
			},

			{
				PieceType: &Lance,
				File:      1,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Knight,
				File:      2,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Silver,
				File:      3,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Gold,
				File:      4,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &King,
				File:      5,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Gold,
				File:      6,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Silver,
				File:      7,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Knight,
				File:      8,
				Rank:      1,
				Player:    Gote,
			},
			{
				PieceType: &Lance,
				File:      9,
				Rank:      1,
				Player:    Gote,
			},

			{
				PieceType: &Lance,
				File:      1,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Knight,
				File:      2,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Silver,
				File:      3,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Gold,
				File:      4,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &King,
				File:      5,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Gold,
				File:      6,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Silver,
				File:      7,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Knight,
				File:      8,
				Rank:      9,
				Player:    Sente,
			},
			{
				PieceType: &Lance,
				File:      9,
				Rank:      9,
				Player:    Sente,
			},
		},
	}

	db.Create(&DefaultPreset)

	weights := EvaluatorWeights{
		Id:   1,
		Name: "Default",

		// Metrics weights
		MATERIAL_WEIGHT:          1.0,
		ATTACK_COUNT_WEIGHT:      1.0,
		PIECE_ADVANCEMENT_WEIGHT: 1.0,
		DEFENDED_PIECES_WEIGHT:   1.0,
		CHECK_WEIGHT:             10.0,
		CHECKMATE_WEIGHT:         99999.0,

		KING_GUARDS_COUNT_WEIGHT:    1.0,
		KING_DEFENCE_RADIUS1_WEIGHT: 2.0,
		KING_DEFENCE_RADIUS2_WEIGHT: 1.0,
		KING_ATTACK_RADIUS1_WEIGHT:  -2.0,
		KING_ATTACK_RADIUS2_WEIGHT:  -1.0,
		KING_FREE_CELLS_WEIGHT:      3.0,

		INVENTORY_MULTIPLIER:     1.5,
		MAX_KING_GUARDS_DISTANCE: 2,
	}

	db.Create(&weights)
}

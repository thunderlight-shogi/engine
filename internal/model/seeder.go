package model

func seed() {
	// Silver+
	SilverPlus := FigureType{
		Name: "Silver+",
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
				HorizontalShift: 1,
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
	KnightPlus := FigureType{
		Name: "Knight+",
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
				HorizontalShift: 1,
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
	LancePlus := FigureType{
		Name: "Lance+",
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
				HorizontalShift: 1,
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
	PawnPlus := FigureType{
		Name: "Pawn+",
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
				HorizontalShift: 1,
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
	RookPlus := FigureType{
		Name: "Rook+",
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
	BishopPlus := FigureType{
		Name: "Bishop+",
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
	King := FigureType{
		Name: "King",
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
	Gold := FigureType{
		Name: "Gold",
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
				HorizontalShift: 1,
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
	Silver := FigureType{
		Name: "Silver",
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
				HorizontalShift: 1,
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
		TurnFigure: &SilverPlus,
	}

	// Knight
	Knight := FigureType{
		Name: "Knight",
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
		TurnFigure: &KnightPlus,
	}

	// Lance
	Lance := FigureType{
		Name: "Lance",
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
		TurnFigure: &LancePlus,
	}

	// Rook
	Rook := FigureType{
		Name: "Rook",
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
		TurnFigure: &RookPlus,
	}

	// Bishop
	Bishop := FigureType{
		Name: "Bishop",
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
		TurnFigure: &BishopPlus,
	}

	// Pawn
	Pawn := FigureType{
		Name: "Pawn",
		Moves: []Move{
			{
				HorizontalShift: 0,
				VerticalShift:   -1,
			},
		},
		TurnFigure: &PawnPlus,
	}

	db.Create([]*FigureType{&King, &Gold, &Silver, &Knight, &Lance, &Rook, &Bishop, &Pawn})

	DefaultPosition := StartingPosition{
		Name: "Default",
		Figures: []StartingPositionFigure{
			{
				FigureType:       Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   3,
				Player:           Gote,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   3,
				Player:           Gote,
			},

			{
				FigureType:       Pawn,
				HorizontalOffset: 1,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 2,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 3,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 4,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 5,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 6,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 7,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 8,
				VerticalOffset:   7,
				Player:           Sente,
			},
			{
				FigureType:       Pawn,
				HorizontalOffset: 9,
				VerticalOffset:   7,
				Player:           Sente,
			},

			{
				FigureType:       Rook,
				HorizontalOffset: 8,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				FigureType:       Rook,
				HorizontalOffset: 2,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				FigureType:       Bishop,
				HorizontalOffset: 2,
				VerticalOffset:   2,
				Player:           Gote,
			},
			{
				FigureType:       Bishop,
				HorizontalOffset: 8,
				VerticalOffset:   8,
				Player:           Sente,
			},

			{
				FigureType:       Lance,
				HorizontalOffset: 1,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Knight,
				HorizontalOffset: 2,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Silver,
				HorizontalOffset: 3,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Gold,
				HorizontalOffset: 4,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       King,
				HorizontalOffset: 5,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Gold,
				HorizontalOffset: 6,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Silver,
				HorizontalOffset: 7,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Knight,
				HorizontalOffset: 8,
				VerticalOffset:   1,
				Player:           Gote,
			},
			{
				FigureType:       Lance,
				HorizontalOffset: 9,
				VerticalOffset:   1,
				Player:           Gote,
			},

			{
				FigureType:       Lance,
				HorizontalOffset: 1,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Knight,
				HorizontalOffset: 2,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Silver,
				HorizontalOffset: 3,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Gold,
				HorizontalOffset: 4,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       King,
				HorizontalOffset: 5,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Gold,
				HorizontalOffset: 6,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Silver,
				HorizontalOffset: 7,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Knight,
				HorizontalOffset: 8,
				VerticalOffset:   9,
				Player:           Sente,
			},
			{
				FigureType:       Lance,
				HorizontalOffset: 9,
				VerticalOffset:   9,
				Player:           Sente,
			},
		},
	}

	db.Create(&DefaultPosition)
}

package model

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Player uint

const (
	Sente Player = iota
	Gote
)

type PieceType struct {
	Id             uint `gorm:"primarykey"`
	Name           string
	Moves          []Move `gorm:"foreignKey:PieceTypeId"`
	PromotePieceId *uint
	PromotePiece   *PieceType
	DemotePiece    *PieceType `gorm:"-:all"`
	Kanji          rune
	ImportantPiece bool
	Cost           uint
}

type StartingPosition struct {
	Id     uint                    `gorm:"primarykey" json:"id"`
	Name   string                  `json:"name"`
	Pieces []StartingPositionPiece `gorm:"foreignKey:StartingPositionId" json:"-"`
}

type Move struct {
	Id              uint `gorm:"primarykey"`
	PieceTypeId     uint
	HorizontalShift int
	VerticalShift   int
}

type StartingPositionPiece struct {
	Id                 uint `gorm:"primarykey"`
	StartingPositionId uint
	PieceTypeId        uint
	PieceType          *PieceType `gorm:"not null"`
	HorizontalOffset   uint
	VerticalOffset     uint
	Player             Player
}

type EvaluatorWeights struct {
	Id   uint `gorm:"primarykey"`
	Name string

	// Metrics weights
	MATERIAL_WEIGHT          float32
	ATTACK_COUNT_WEIGHT      float32
	PIECE_ADVANCEMENT_WEIGHT float32
	DEFENDED_PIECES_WEIGHT   float32
	CHECK_WEIGHT             float32
	CHECKMATE_WEIGHT         float32

	KING_GUARDS_COUNT_WEIGHT    float32
	KING_DEFENCE_RADIUS1_WEIGHT float32
	KING_DEFENCE_RADIUS2_WEIGHT float32
	KING_ATTACK_RADIUS1_WEIGHT  float32
	KING_ATTACK_RADIUS2_WEIGHT  float32
	KING_FREE_CELLS_WEIGHT      float32

	// Extra constants
	INVENTORY_MULTIPLIER     float32
	MAX_KING_GUARDS_DISTANCE uint
}

var db *gorm.DB

func init() {
	link, ok := os.LookupEnv("MYSQL_LINK")
	if !ok {
		panic("MYSQL_LINK environment variable is undefined")
	}
	var err error
	db, err = gorm.Open(mysql.Open(link), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&StartingPosition{}, &StartingPositionPiece{}, &PieceType{}, &Move{}, &EvaluatorWeights{})

	result := db.Find(&PieceType{})
	if result.RowsAffected == 0 {
		seed()
	}
}

func GetDB() *gorm.DB {
	return db
}

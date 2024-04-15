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
	Id             uint       `gorm:"primarykey" json:"id"`
	Name           string     `json:"name"`
	Moves          []Move     `gorm:"foreignKey:PieceTypeId" json:"moves"`
	PromotePieceId *uint      `json:"-"`
	PromotePiece   *PieceType `json:"promote_piece"`
	DemotePiece    *PieceType `gorm:"-:all" json:"-"`
	Kanji          rune       `json:"kanji"`
	ImportantPiece bool       `json:"important_piece"`
	Cost           uint       `json:"cost"`
}

type Preset struct {
	Id     uint          `gorm:"primarykey" json:"id"`
	Name   string        `json:"name"`
	Pieces []PresetPiece `gorm:"foreignKey:PresetId" json:"pieces"`
}

type Move struct {
	Id          uint `gorm:"primarykey" json:"id"`
	PieceTypeId uint `json:"-"`
	FileShift   int  `json:"file_shift"`
	RankShift   int  `json:"rank_shift"`
}

type PresetPiece struct {
	Id          uint       `gorm:"primarykey" json:"id"`
	PresetId    uint       `json:"-"`
	PieceTypeId uint       `json:"-"`
	PieceType   *PieceType `gorm:"not null" json:"piece_type"`
	File        uint       `json:"file"`
	Rank        uint       `json:"rank"`
	Player      Player     `json:"player"`
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
	db.AutoMigrate(&Preset{}, &PresetPiece{}, &PieceType{}, &Move{}, &EvaluatorWeights{})

	result := db.Find(&PieceType{})
	if result.RowsAffected == 0 {
		seed()
	}
}

func GetDB() *gorm.DB {
	return db
}

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
	TurnPieceId    *uint
	TurnPiece      *PieceType
	Kanji          rune
	ImportantPiece bool
}

type StartingPosition struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Pieces []StartingPositionPiece `gorm:"foreignKey:StartingPositionId"`
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
	PieceType          PieceType
	HorizontalOffset   uint
	VerticalOffset     uint
	Player             Player
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
	db.AutoMigrate(&StartingPosition{}, &StartingPositionPiece{}, &PieceType{}, &Move{})

	result := db.Find(&PieceType{})
	if result.RowsAffected == 0 {
		seed()
	}
}

func GetDB() *gorm.DB {
	return db
}

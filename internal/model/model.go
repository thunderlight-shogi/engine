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

type FigureType struct {
	Id           uint `gorm:"primarykey"`
	Name         string
	Moves        []Move `gorm:"foreignKey:FigureTypeId"`
	TurnFigureId *uint
	TurnFigure   *FigureType
	Kanji        rune
}

type StartingPosition struct {
	Id      uint `gorm:"primarykey"`
	Name    string
	Figures []StartingPositionFigure `gorm:"foreignKey:StartingPositionId"`
}

type Move struct {
	Id              uint `gorm:"primarykey"`
	FigureTypeId    uint
	HorizontalShift int
	VerticalShift   int
}

type StartingPositionFigure struct {
	Id                 uint `gorm:"primarykey"`
	StartingPositionId uint
	FigureTypeId       uint
	FigureType         FigureType
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
	db.AutoMigrate(&StartingPosition{}, &StartingPositionFigure{}, &FigureType{}, &Move{})

	result := db.Find(&FigureType{})
	if result.RowsAffected == 0 {
		seed()
	}
}

func GetDB() *gorm.DB {
	return db
}

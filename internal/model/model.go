package model

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FigureType struct {
	Id           uint `gorm:"primarykey"`
	Name         string
	Moves        []Move `gorm:"foreignKey:FigureTypeId"`
	TurnFigureId *uint
	TurnFigure   *FigureType
}

type StartingPosition struct {
	Id      uint `gorm:"primarykey"`
	Name    string
	Figures []StartingPositionFigure `gorm:"foreignKey:StartingPositionId"`
}

type Move struct {
	Id           uint `gorm:"primarykey"`
	FigureTypeId uint
	ShiftX       int // Horizontal shift
	ShiftY       int // Vertical shift
}

type StartingPositionFigure struct {
	Id                 uint `gorm:"primarykey"`
	StartingPositionId uint
	FigureTypeId       uint
	FigureType         FigureType
	PositionX          uint // Horizontal offset
	PositionY          uint // Vertical offset
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
		// Seed database
	}
}

func GetDB() *gorm.DB {
	return db
}

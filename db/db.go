package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Link struct {
	gorm.Model
	Path string `gorm:"unique;not null"`
	Link string `gorm:"not null"`
	Pin  uint
	Uid  uint
}

func DBsetup() (err error) {
	dsn := os.Getenv("APP_DB_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	db.AutoMigrate(&Link{})
	return err
}

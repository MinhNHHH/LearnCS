package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *Database

func Init(dsn string) *Database {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = &Database{db}
	return DB
}

func GetDB() *Database {
	return DB
}

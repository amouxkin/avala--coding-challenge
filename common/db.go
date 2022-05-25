package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&HexCounter{})

	if err != nil {
		panic("Migration Failed")
	}
}

// GetDb connects to the database and returns the orm.
func GetDb() *gorm.DB {
	dsn, ok := os.LookupEnv("DSN")

	if !ok {
		panic("DSN not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

type HexCounter struct {
	Hex   int `gorm:"unique;default:0;not null"`
	Count int `gorm:"not null;default:0"`
}

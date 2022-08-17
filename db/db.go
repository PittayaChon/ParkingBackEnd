package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Park struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	LotId        string `json:"lot_id"`
	Licenseplate string `json:"licenseplate"`
	Status       int    `json:"status"`
	Reserveable  bool   `json:"reserveable"`
	Floor        string `json:"floor"`
}

func DB() *gorm.DB {
	env := LoadENV()

	psqlInfo := "postgresql://" + env.POSTGRES_USER + ":" + env.POSTGRES_PASSWORD + "@" + env.POSTGRES_HOST + ":" + env.POSTGRES_PORT + "/" + env.POSTGRES_DB

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Park{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}

package db

import (
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/logger"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(log *logger.Logger) *gorm.DB {
	dbPath := "db/contracts.db"

	if err := os.MkdirAll("db", 0755); err != nil {
		log.Fatalf(ErrFailedToCreateDBDirectory, err)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf(ErrFailedToCreateDBFile, err)
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf(ErrFailedToInitializeDB, err)
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Person{},
		&models.Contract{},
		&models.ContractPerson{},
	); err != nil {
		log.Fatalf(ErrFailedToMigrateDB, err)
	}

	return db
}

package db

import (
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/logger"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(log *logger.Logger) *gorm.DB {
	dbPath := "sqlitedb/contracts.db"

	if err := os.MkdirAll("sqlitedb", 0755); err != nil {
		log.Fatalf(ErrFailedToCreateDBDirectory, err)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf(ErrFailedToCreateDBFile, err)
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath+"?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatalf(ErrFailedToInitializeDB, err)
	}

	if err = db.AutoMigrate(
		&models.Person{},
		&models.User{},
		&models.Contract{},
		&models.ContractPerson{},
	); err != nil {
		log.Fatalf(ErrFailedToMigrateDB, err)
	}

	return db
}

package db

import "errors"

var (
	ErrFailedToCreateDBDirectory = errors.New("failed to create database directory")
	ErrFailedToCreateDBFile      = errors.New("failed to create db file")
	ErrFailedToInitializeDB      = errors.New("failed to initialize database")
	ErrFailedToMigrateDB         = errors.New("failed to migrate database")
)

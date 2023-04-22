package services

import (
	"log"

	"github.com/h00s/bitbox/api/db"
)

type Services struct {
	DB     *db.Database
	Logger *log.Logger
}

func NewServices(db *db.Database, logger *log.Logger) *Services {
	return &Services{
		DB:     db,
		Logger: logger,
	}
}

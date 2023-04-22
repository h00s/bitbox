package services

import (
	"log"
)

type Services struct {
	Logger *log.Logger
}

func NewServices(logger *log.Logger) *Services {
	return &Services{
		Logger: logger,
	}
}

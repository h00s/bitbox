package db

import (
	"github.com/go-raptor/connector/bun/postgres"
	"github.com/h00s/bitbox/db/migrate"
)

func Migrations() postgres.Migrations {
	return postgres.Migrations{
		"20250114101729": &migrate.CreatePaste{},
	}
}

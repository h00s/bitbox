package migrate

import (
	"context"

	"github.com/h00s/bitbox/app/models"
	"github.com/uptrace/bun"
)

func AddPaste(db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model(&models.Paste{}).
		IfNotExists().
		Exec(context.Background())

	return err
}

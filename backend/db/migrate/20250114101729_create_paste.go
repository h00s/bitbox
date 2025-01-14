package migrate

import (
	"context"

	"github.com/h00s/bitbox/app/models"
	"github.com/uptrace/bun"
)

type CreatePaste struct{}

func (m CreatePaste) Name() string {
	return "create_paste_table"
}

func (m CreatePaste) Up(db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model(&models.Paste{}).
		IfNotExists().
		Exec(context.Background())
	return err
}

func (m CreatePaste) Down(db *bun.DB) error {
	_, err := db.NewDropTable().
		Model(&models.Paste{}).
		IfExists().
		Cascade().
		Exec(context.Background())
	return err
}

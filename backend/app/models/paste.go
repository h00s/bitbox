package models

import (
	"time"
)

const validChars = "bcdfghmnprstvz23456789"

type Paste struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicPaste struct {
	ID string
	Paste
}

func (p *Paste) ToPublicPaste() PublicPaste {
	return PublicPaste{
		ID:    shortURIfromID(p.ID),
		Paste: *p,
	}
}

type Pastes []Paste
type PublicPastes []PublicPaste

func shortURIfromID(id uint) string {
	uri := ""
	for id > 0 {
		uri = string(validChars[id%uint(len(validChars))]) + uri
		id = id / uint(len(validChars))
	}
	return uri
}

/*func idFromShortURI(uri string) uint {
	id := uint(0)
	for _, c := range uri {
		id = id*uint(len(validChars)) + uint(strings.Index(validChars, string(c)))
	}
	return id
}*/

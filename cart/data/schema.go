package data

import "time"

const Schema = `
	CREATE TABLE IF NOT EXISTS wishlist (
		uid			varchar(80) PRIMARY KEY,
		"user"		varchar(80) REFERENCES users ON DELETE CASCADE,
		laptop		varchar(80) REFERENCES laptops ON DELETE CASCADE,

		created_at	TIMESTAMP NOT NULL DEFAULT NOW()
	)
`

type WishlistItem struct {
	Uid       string    `json:"uid"`
	User      string    `json:"user"`
	Laptop    string    `json:"laptop"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

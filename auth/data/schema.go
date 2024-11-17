package data

import "time"

const Schema = `
	CREATE TABLE IF NOT EXISTS users (
		uid			varchar(80) PRIMARY KEY,
		email		varchar(24) NULL UNIQUE,
		full_name	varchar(80) NULL,
		password	varchar(255) NULL,
		logged_at	TIMESTAMP NULL,

		ip			varchar(15) NULL,
		user_agent	varchar(255) NULL,

		created_at	TIMESTAMP NOT NULL DEFAULT NOW()
	)
`

type User struct {
	Uid      string     `json:"uid"`
	Email    *string    `json:"email"`
	FullName *string    `db:"full_name" json:"full_name"`
	Password *string    `json:"password,omitempty"`
	LoggedAt *time.Time `db:"logged_at" json:"logged_at"`

	Ip        *string `json:"ip"`
	UserAgent *string `db:"user_agent" json:"user_agent"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

package user

import (
	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	db *sqlx.DB
}

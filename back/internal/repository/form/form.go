package anketa

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoAnketa struct {
	db *sqlx.DB
}

func (repo RepoAnketa) Create(ctx context.Context, user models.UserAnketa) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowxContext(ctx, `insert into users (about, sphere, photo) values ($1, $2, $3) returning `)
}

func InitAnketaRepository(db *sqlx.DB) repository.AnketaRepo {
	return RepoAnketa{db: db}
}

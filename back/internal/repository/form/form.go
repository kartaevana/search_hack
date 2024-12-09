package form

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoForm struct {
	db *sqlx.DB
}

func (repo RepoForm) Create(ctx context.Context, form models.FormCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowxContext(ctx, `insert into form (id_user, about, photo, sphere) values ($1, $2, $3, $4) returning id`,
		form.UserID, form.About, form.Photo, form.Sphere)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	err = transaction.Commit()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	return id, nil
}

func InitFormRepository(db *sqlx.DB) repository.FormRepo {
	return RepoForm{db: db}
}

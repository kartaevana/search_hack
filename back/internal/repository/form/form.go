package form

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepoForm struct {
	db *sqlx.DB
}

func (repo RepoForm) Get(ctx context.Context, id int) (*models.Form, error) {
	var form models.Form
	row := repo.db.QueryRowContext(ctx, `select id, id_user, about, photo, sphere from form where id = $1`, id)
	err := row.Scan(&form.ID, &form.ID_User, &form.About, &form.Photo, &form.Sphere)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("form not found")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	row = repo.db.QueryRowContext(ctx, `select name, sur_name, tg from users where id = $1`, form.ID_User)
	err = row.Scan(&form.Name, &form.Surname, &form.Tg)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error fetching user data: %w", err)
	}
	fmt.Println(form)
	return &form, nil
}

func (repo RepoForm) Create(ctx context.Context, form models.FormCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowxContext(ctx, `insert into form (id_user, about, photo, sphere) values ($1, $2, $3, $4) returning id`,
		form.ID_User, form.About, form.Photo, form.Sphere)
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

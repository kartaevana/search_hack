package user

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	db *sqlx.DB
}

func (repo RepoUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	//открываем транзакцию, чтобы изменить что-то в таблице
	row := transaction.QueryRowxContext(ctx, `insert into users(name, sur_name, email, hashed_password) values ($1, $2, $3, $4) returning id;`,
		user.Name, user.Surname, user.Email, user.PWD)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			// если при сканировании произошла ошибка,
			//то делаем откат до состояния изменения
			return 0, rbErr
		}
		return 0, err
	}
	if err = transaction.Commit(); err != nil {
		// сохраняем изменения
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	return id, nil
}

func (repo RepoUser) Get(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	row := repo.db.QueryRowxContext(ctx, `SELECT name, sur_name, email FROM users WHERE id = $1`, id)
	err := row.Scan(&user.Name, &user.Surname, &user.Email)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepoUser{db: db}
}

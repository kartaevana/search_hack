package approve

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoApprove struct {
	db *sqlx.DB
}

func (repo RepoApprove) Create(ctx context.Context, approve models.ApproveCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowxContext(ctx, `insert into approve (team_id, form_id) values ($1, $2) returning id`,
		approve.ID_Team, approve.ID_Form)
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

func InitApproveRepository(db *sqlx.DB) repository.ApproveRepo {
	return RepoApprove{db: db}
}

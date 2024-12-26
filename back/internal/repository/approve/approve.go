package approve

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepoApprove struct {
	db *sqlx.DB
}

func (repo RepoApprove) GetAll(ctx context.Context) ([]*models.Approve, error) {
	var approves []*models.Approve
	rows, err := repo.db.QueryContext(ctx, `
        SELECT approve.id, approve.team_id, approve.form_id
        FROM approve`)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var approve models.Approve
		err := rows.Scan(&approve.ID, &approve.ID_Team, &approve.ID_Form)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		approves = append(approves, &approve)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return approves, nil
}

func (repo RepoApprove) Accept(ctx context.Context, approve models.Approve) error {
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = transaction.ExecContext(ctx, `INSERT INTO team_user (team_id, user_id)
SELECT $1, form.id_user
FROM form
WHERE form.id = $2`, approve.ID_Team, approve.ID_Form)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	_, err = transaction.ExecContext(
		ctx,
		`delete from approve where approve.id=$1`,
		approve.ID)
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}
	err = transaction.Commit()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return nil
}

func (repo RepoApprove) Reject(ctx context.Context, approve models.Approve) error {
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = transaction.ExecContext(ctx, `delete from approve where id=$1`, approve.ID)
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}
	err = transaction.Commit()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return nil
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

package team

import (
	"back/internal/models"
	"back/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepoTeam struct {
	db *sqlx.DB
}

func (repo RepoTeam) GetTeam(ctx context.Context, id int) (*models.Team, error) {
	var team models.Team
	row := repo.db.QueryRowContext(ctx, `select id, name, about, id_kap from team where id = $1`, id)
	err := row.Scan(&team.ID, &team.Name, &team.About, &team.ID_Kap)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("form not found")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	rows, err := repo.db.QueryContext(ctx, `SELECT user_id from team_user where team_id=$1`, id)
	for rows.Next() {
		var id_user int
		err = rows.Scan(&id_user)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("form not found")
			}
			return nil, fmt.Errorf("database error: %w", err)
		}
		var user models.User
		row = repo.db.QueryRowContext(ctx, `select id, email, name, sur_name, tg from users where id = $1`, id_user)
		err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.Tg)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("form not found")
			}
			return nil, fmt.Errorf("database error: %w", err)
		}
		team.Members = append(team.Members, user)
	}
	return &team, nil
}

func (repo RepoTeam) Create(ctx context.Context, team models.TeamCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	row := transaction.QueryRowxContext(ctx, `insert into team (about, name, id_kap) values ($1, $2, $3) returning id`,
		team.About, team.Name, team.ID_Kap)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}

	row = transaction.QueryRowxContext(ctx, `insert into team_user (user_id, team_id) values ($1, $2)`,
		team.ID_Kap, id)
	err = transaction.Commit()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}
	return id, nil
}

func (repo RepoTeam) AddUserTeam(ctx context.Context, id_user int, id_team int) error {

	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_ = transaction.QueryRowxContext(ctx, `insert into team_user (user_id, team_id) values ($1, $2)`,
		id_user, id_team)
	err = transaction.Commit()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return nil
}

func InitTeamRepository(db *sqlx.DB) repository.TeamRepo {
	return RepoTeam{db: db}
}

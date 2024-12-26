package repository

import (
	"back/internal/models"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, email string) (int, string, error)
}

type FormRepo interface {
	Create(ctx context.Context, form models.FormCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Form, error)
	GetAll(ctx context.Context) ([]*models.Form, error)
}

type TeamRepo interface {
	Create(ctx context.Context, team models.TeamCreate) (int, error)
	GetTeam(ctx context.Context, id int) (*models.Team, error)
	AddUserTeam(ctx context.Context, id_user int, id_team int) error
}

type ApproveRepo interface {
	Create(ctx context.Context, approve models.ApproveCreate) (int, error)
	Reject(ctx context.Context, approve models.Approve) error
	GetAll(ctx context.Context) ([]*models.Approve, error)
}

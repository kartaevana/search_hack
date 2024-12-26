package service

import (
	"back/internal/models"
	"context"
)

type UserServ interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, user models.UserLogin) (int, error)
}

type FormServ interface {
	Create(ctx context.Context, form models.FormCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Form, error)
	GetAll(ctx context.Context) ([]*models.Form, error)
}

type ApproveServ interface {
	Create(ctx context.Context, approve models.ApproveCreate) (int, error)
	Reject(ctx context.Context, approve models.Approve) error
	GetAll(ctx context.Context) ([]*models.Approve, error)
	Accept(ctx context.Context, approve models.Approve) error
}

type TeamServ interface {
	Create(ctx context.Context, team models.TeamCreate) (int, error)
	GetTeam(ctx context.Context, id int) (*models.Team, error)
	AddUserTeam(ctx context.Context, id_user int, id_team int) error
}

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
}

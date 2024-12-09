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
}

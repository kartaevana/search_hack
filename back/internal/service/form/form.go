package form

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/service"
	"context"
)

type ServForm struct {
	FormRepo repository.FormRepo
}

func (serv ServForm) Create(ctx context.Context, form models.FormCreate) (int, error) {
	id, err := serv.FormRepo.Create(ctx, form)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InitFormService(formRepo repository.FormRepo) service.FormServ {
	return &ServForm{FormRepo: formRepo}
}

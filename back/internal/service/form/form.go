package form

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/service"
	"back/pkg/log"
	"context"
	"fmt"
)

type ServForm struct {
	FormRepo repository.FormRepo
	log      *log.Logs
}

func (serv ServForm) Get(ctx context.Context, id int) (*models.Form, error) {
	form, err := serv.FormRepo.Get(ctx, id)
	if err != nil {
		serv.log.Error(err.Error())
		return nil, err
	}
	serv.log.Info(fmt.Sprintf("get form: %v", id))
	return form, nil
}

func (serv ServForm) Create(ctx context.Context, form models.FormCreate) (int, error) {
	id, err := serv.FormRepo.Create(ctx, form)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InitFormService(formRepo repository.FormRepo, log *log.Logs) service.FormServ {
	return &ServForm{FormRepo: formRepo, log: log}
}

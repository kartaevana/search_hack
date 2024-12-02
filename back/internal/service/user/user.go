package user

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/service"
	"back/pkg/log"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type ServUser struct {
	UserRepo repository.UserRepo
	log      *log.Logs
}

func (serv ServUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PWD), 10)
	if err != nil {
		serv.log.Error(err.Error())
		return 0, err
	}
	newUser := models.UserCreate{
		UserBase: user.UserBase,
		PWD:      string(hashedPassword),
	}
	id, err := serv.UserRepo.Create(ctx, newUser)
	if err != nil {
		serv.log.Error(err.Error())
		return 0, err
	}
	serv.log.Info(fmt.Sprintf("create user %v", id))
	return id, nil
}

func (serv ServUser) Get(ctx context.Context, id int) (*models.User, error) {
	user, err := serv.UserRepo.Get(ctx, id)
	if err != nil {
		serv.log.Error(err.Error())
		return nil, err
	}
	serv.log.Info(fmt.Sprintf("get user: %v", user.ID))
	return user, nil
}

func InitUserService(userRepo repository.UserRepo, log *log.Logs) service.UserServ {
	return &ServUser{UserRepo: userRepo, log: log}
}

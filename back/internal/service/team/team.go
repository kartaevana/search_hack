package team

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/service"
	"back/pkg/log"
	"context"
	"fmt"
)

type ServTeam struct {
	TeamRepo repository.TeamRepo
	log      *log.Logs
}

func (serv ServTeam) AddUserTeam(ctx context.Context, id_user int, id_team int) error {
	err := serv.TeamRepo.AddUserTeam(ctx, id_user, id_team)
	if err != nil {
		serv.log.Error(err.Error())
		return err
	}
	return nil
}

func (serv ServTeam) GetTeam(ctx context.Context, id int) (*models.Team, error) {
	team, err := serv.TeamRepo.GetTeam(ctx, id)
	if err != nil {
		serv.log.Error(err.Error())
		return nil, err
	}
	serv.log.Info(fmt.Sprintf("get team: %v", id))
	return team, nil
}

func (serv ServTeam) Create(ctx context.Context, team models.TeamCreate) (int, error) {
	id, err := serv.TeamRepo.Create(ctx, team)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InitTeamService(teamRepo repository.TeamRepo, log *log.Logs) service.TeamServ {
	return &ServTeam{
		TeamRepo: teamRepo,
		log:      log,
	}
}

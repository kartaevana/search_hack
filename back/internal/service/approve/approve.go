package approve

import (
	"back/internal/models"
	"back/internal/repository"
	"back/internal/service"
	"back/pkg/log"
	"context"
	"fmt"
)

type ServApprove struct {
	ApproveRepo repository.ApproveRepo
	log         *log.Logs
}

func (serv ServApprove) GetAll(ctx context.Context) ([]*models.Approve, error) {
	approve, err := serv.ApproveRepo.GetAll(ctx)
	if err != nil {
		serv.log.Error(err.Error())
		return nil, err
	}
	serv.log.Info(fmt.Sprintf("Got all %v", approve))
	return approve, err
}

func (serv ServApprove) Reject(ctx context.Context, approve models.Approve) error {
	err := serv.ApproveRepo.Reject(ctx, approve)
	if err != nil {
		return err
	}
	serv.log.Info(fmt.Sprintf("Rejected %v", approve.ID_Form))
	return nil
}

func (serv ServApprove) Create(ctx context.Context, approve models.ApproveCreate) (int, error) {
	id, err := serv.ApproveRepo.Create(ctx, approve)
	if err != nil {
		return 0, err
	}
	serv.log.Info(fmt.Sprintf("Creating approve ID %v", id))
	return id, nil
}

func InitApproveService(approveRepo repository.ApproveRepo, log *log.Logs) service.ApproveServ {
	return &ServApprove{ApproveRepo: approveRepo, log: log}
}

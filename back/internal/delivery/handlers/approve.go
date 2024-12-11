package handlers

import (
	"back/internal/models"
	"back/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApproveHandler struct {
	service service.ApproveServ
}

func InitApproveHandler(service service.ApproveServ) ApproveHandler {
	return ApproveHandler{
		service: service,
	}
}

// @Summary Create approve
// @Tags approve
// @Accept  json
// @Produce  json
// @Param data body models.ApproveCreate true "approve create"
// @Success 200 {object} int "Successfully created approve"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /approve/create [post]
func (handler ApproveHandler) CreateApprove(g *gin.Context) {
	var newApprove models.ApproveCreate

	if err := g.ShouldBindJSON(&newApprove); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Create(ctx, newApprove)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

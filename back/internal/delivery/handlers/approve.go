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

// @Summary Accept approve
// @Tags approve
// @Accept  json
// @Produce  json
// @Param data body models.Approve true "approve accepted"
// @Success 200 {object} int "Successfully accepted approve"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /approve/accept [put]
func (handler ApproveHandler) AcceptApprove(g *gin.Context) {
	var newApprove models.Approve

	if err := g.ShouldBindJSON(&newApprove); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.Accept(ctx, newApprove)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"accepted": "success"})
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

// @Summary Reject approve
// @Tags approve
// @Accept  json
// @Produce  json
// @Param data body models.Approve true "approve rejected"
// @Success 200 {object} int "Successfully rejected approve"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /approve/reject [post]
func (handler ApproveHandler) RejectApprove(g *gin.Context) {
	var newApprove models.Approve

	if err := g.ShouldBindJSON(&newApprove); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.Reject(ctx, newApprove)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"reject": "success"})
}

// @Summary Get all approves
// @Tags approve
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Approve "Successfully get all approves"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /approve/all [get]
func (handler ApproveHandler) GetAllApprove(g *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	approve, err := handler.service.GetAll(ctx)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"approve": approve})
}

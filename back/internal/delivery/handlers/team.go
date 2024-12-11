package handlers

import (
	"back/internal/models"
	"back/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type TeamHandler struct {
	service service.TeamServ
}

func InitTeamHandler(service service.TeamServ) TeamHandler {
	return TeamHandler{
		service: service,
	}
}

// @Summary Create team
// @Tags team
// @Accept  json
// @Produce  json
// @Param data body models.TeamCreate true "team create"
// @Success 200 {object} int "Successfully created team"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /team/create [post]
func (handler TeamHandler) CreateTeam(g *gin.Context) {
	var newTeam models.TeamCreate

	if err := g.ShouldBindJSON(&newTeam); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Create(ctx, newTeam)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get team
// @Tags team
// @Accept  json
// @Produce  json
// @Param id query int true "team get"
// @Success 200 {object} int "Successfully get team"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /team/{id} [get]
func (handler TeamHandler) GetTeam(g *gin.Context) {
	id := g.Query("id")
	aid, err := strconv.Atoi(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	team, err := handler.service.GetTeam(ctx, aid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"team": team})

}

// @Summary AddUserTeam team
// @Tags team
// @Accept  json
// @Produce  json
// @Param id_team query int true "add user in team"
// @Param id_user query int true "add user in team"
// @Success 200 {object} int "Successfully add user in team"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /team/add/{id_team}/{id_user} [post]
func (handler TeamHandler) AddUserTeam(g *gin.Context) {
	id_team := g.Query("id_team")
	aid_team, err := strconv.Atoi(id_team)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id_user := g.Query("id_user")
	aid_user, err := strconv.Atoi(id_user)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.AddUserTeam(ctx, aid_user, aid_team)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"add": "success"})

}

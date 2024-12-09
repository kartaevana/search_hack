package handlers

import (
	"back/internal/models"
	"back/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FormHandler struct {
	service service.FormServ
}

func InitFormHandler(service service.FormServ) FormHandler {
	return FormHandler{
		service: service,
	}
}

// @Summary Create form
// @Tags form
// @Accept  json
// @Produce  json
// @Param data body models.FormCreate true "form create"
// @Success 200 {object} int "Successfully created form"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /form/create [post]
func (handler FormHandler) CreateForm(g *gin.Context) {
	var newForm models.FormCreate

	if err := g.ShouldBindJSON(&newForm); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Create(ctx, newForm)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

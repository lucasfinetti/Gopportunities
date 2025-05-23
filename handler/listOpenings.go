package handler

import (
	"net/http"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := Db.Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	SendSuccess(ctx, "list-openings", openings)
}

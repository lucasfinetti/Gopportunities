package handler

import (
	"net/http"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create candidate
// @Description Create a new job candidate
// @Tags Candidates
// @Accept json
// @Produce json
// @Param request body CreateCandidateRequest true "Request body"
// @Success 200 {object} CreateCandidateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /candidate [post]
func CreateCandidateHandler(ctx *gin.Context) {
	request := CreateCandidateRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	candidate := schemas.Candidate{
		Name:      request.Name,
		BirthDate: request.BirthDate,
	}

	if err := db.Create(&candidate).Error; err != nil {
		logger.Errorf("error creating candidate: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating candidate on database")
		return
	}

	sendSuccess(ctx, "create-candidate", candidate)
}

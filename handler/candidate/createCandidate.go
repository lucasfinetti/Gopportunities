package candidate

import (
	"net/http"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/handler"
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
// @Success 200 {object} handler.CreateCandidateResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /candidate [post]
func CreateCandidateHandler(ctx *gin.Context) {
	request := CreateCandidateRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	candidate := schemas.Candidate{
		Name:      request.Name,
		BirthDate: request.BirthDate,
	}

	if err := handler.Db.Create(&candidate).Error; err != nil {
		handler.Logger.Errorf("error creating candidate: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error creating candidate on database")
		return
	}

	handler.SendSuccess(ctx, "create-candidate", candidate)
}

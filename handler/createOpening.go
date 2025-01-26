package handler

import (
	"net/http"

	"dev.azure.com/lucasfinetti/Finetti/_git/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
		Closed:   *request.Closed,
	}

	if err := Db.Create(&opening).Error; err != nil {
		Logger.Errorf("error creating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	SendSuccess(ctx, "create-opening", opening)
}

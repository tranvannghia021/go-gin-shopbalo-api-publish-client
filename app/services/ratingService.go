package services

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/helpers"
	"go-api-publish/app/models"
	"go-api-publish/app/repositories"
	"net/http"
)

type RatingServiceInterface interface {
	GetAll(ctx *gin.Context)
}
type RatingService struct {
	ratingRepo repositories.RatingRepositoryInterface
}

func NewRating(ratingRepo repositories.RatingRepositoryInterface) RatingServiceInterface {
	return &RatingService{
		ratingRepo: ratingRepo,
	}
}
func (r *RatingService) GetAll(ctx *gin.Context) {
	var rating []models.Rating
	err := r.ratingRepo.GetAll(&rating)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ApiResponseError(nil, err, 401))
		return
	}
	ctx.JSON(http.StatusOK, helpers.ApiResponse(rating, "list rating", 200))
}

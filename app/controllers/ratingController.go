package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/services"
)

type RatingControllerInterface interface {
	GetAll(ctx *gin.Context)
}
type RatingController struct {
	ratingService services.RatingServiceInterface
}

func NewRating(RatingService services.RatingServiceInterface) RatingControllerInterface {
	return &RatingController{
		ratingService: RatingService,
	}

}
func (r *RatingController) GetAll(ctx *gin.Context) {
	r.ratingService.GetAll(ctx)
}

package services

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/helpers"
	"go-api-publish/app/models"
	"go-api-publish/app/repositories"
	"net/http"
)

type SliderServiceInterface interface {
	GetAll(ctx *gin.Context)
}
type SliderService struct {
	sliderRepo repositories.SliderRepositoryInterface
}

func NewSlider(sliderRepo repositories.SliderRepositoryInterface) SliderServiceInterface {
	return &SliderService{
		sliderRepo: sliderRepo,
	}
}
func (s *SliderService) GetAll(ctx *gin.Context) {
	var slider []models.Slider
	err := s.sliderRepo.GetAll(ctx, &slider)
	if err != nil {
		res := helpers.ApiResponseError(slider, err, 401)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.ApiResponse(slider, "list slider", 200)
	ctx.JSON(http.StatusOK, res)
}

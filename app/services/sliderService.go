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

		ctx.JSON(http.StatusBadRequest, helpers.ApiResponseError(slider, err, 401))
		return
	}
	ctx.JSON(http.StatusOK, helpers.ApiResponse(slider, "list slider", 200))

}

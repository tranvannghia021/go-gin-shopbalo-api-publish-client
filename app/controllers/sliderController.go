package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/services"
)

type SliderControllerInterface interface {
	GetAll(ctx *gin.Context)
}
type SliderController struct {
	sliderService services.SliderServiceInterface
}

func NewSlider(sliderService services.SliderServiceInterface) SliderControllerInterface {
	return &SliderController{
		sliderService: sliderService,
	}
}
func (s *SliderController) GetAll(ctx *gin.Context) {
	s.sliderService.GetAll(ctx)
}

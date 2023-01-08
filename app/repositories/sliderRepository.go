package repositories

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/models"
	"go-api-publish/config"
	"gorm.io/gorm"
)

type SliderRepositoryInterface interface {
	GetAll(ctx *gin.Context, Slider *[]models.Slider) error
}
type SliderRepository struct {
	db *gorm.DB
}

func NewSlider() SliderRepositoryInterface {
	return &SliderRepository{
		db: config.InitDb(),
	}
}
func (s SliderRepository) GetAll(ctx *gin.Context, Slider *[]models.Slider) error {
	return s.db.Find(Slider).Error
}

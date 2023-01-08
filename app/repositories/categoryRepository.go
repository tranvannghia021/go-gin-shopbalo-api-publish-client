package repositories

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/models"
	"go-api-publish/config"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetAll(ctx *gin.Context, category *[]models.Category) error
}
type CategoryRepository struct {
	db *gorm.DB
}

func NewCategory() CategoryRepositoryInterface {
	return &CategoryRepository{
		db: config.InitDb(),
	}
}
func (CategoryRepo *CategoryRepository) GetAll(ctx *gin.Context, category *[]models.Category) error {
	return CategoryRepo.db.Find(category).Error
}

package repositories

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/models"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetAll(ctx *gin.Context, category *[]models.Category) error
}
type CategoryRepository struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{
		db: db,
	}
}
func (CategoryRepo *CategoryRepository) GetAll(ctx *gin.Context, category *[]models.Category) error {
	return CategoryRepo.db.Table("categories").Find(category).Error
}

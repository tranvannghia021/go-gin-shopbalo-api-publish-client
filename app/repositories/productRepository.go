package repositories

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/models"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetAll(ctx *gin.Context, product *[]models.Product, sort string) error
}
type ProductRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}
func (p *ProductRepository) GetAll(ctx *gin.Context, product *[]models.Product, sort string) error {
	return p.db.Order("id " + sort).Preload("Category").
		Preload("ProductDetail").Preload("Rating").Find(product).Error

}

package repositories

import (
	"go-api-publish/app/models"
	"gorm.io/gorm"
)

type RatingRepositoryInterface interface {
	//FindByProductId(data *services.RatingRes, number int32) error
	GetAll(Rating *[]models.Rating) error
}
type RatingRepository struct {
	db *gorm.DB
}

func NewRating(db *gorm.DB) RatingRepositoryInterface {
	return &RatingRepository{
		db: db,
	}
}

//	func (r *RatingRepository) FindByProductId(data *services.RatingRes, number int32) error {
//		return r.db.Select("avg(point) as point_avg ,product_id").Where("product_id = ?", number).Find(data).Error
//	}
func (r *RatingRepository) GetAll(Rating *[]models.Rating) error {
	return r.db.Table("ratings").Find(Rating).Error
}

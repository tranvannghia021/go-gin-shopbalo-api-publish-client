package routers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/controllers"
	"go-api-publish/app/repositories"
	"go-api-publish/app/services"
	"go-api-publish/config"
)

var dbGorm = config.InitDb()
var (
	CategoryController = controllers.NewCategory(services.NewCategory(repositories.NewCategory(dbGorm)))
	SliderController   = controllers.NewSlider(services.NewSlider(repositories.NewSlider(dbGorm)))
	ProductController  = controllers.NewProduct(services.NewProduct(repositories.NewProduct(dbGorm)))
	RatingController   = controllers.NewRating(services.NewRating(repositories.NewRating(dbGorm)))
)

func ApiRouter(group *gin.RouterGroup) {
	cateGroup := group.Group("category")
	{

		cateGroup.GET("", CategoryController.GetAll)

	}
	sliderGroup := group.Group("slider")
	{
		sliderGroup.GET("", SliderController.GetAll)

	}
	productGroup := group.Group("product")
	{
		productGroup.GET("", ProductController.GetAll)

	}
	ratingGroup := group.Group("rating")
	{
		ratingGroup.GET("", RatingController.GetAll)

	}

}

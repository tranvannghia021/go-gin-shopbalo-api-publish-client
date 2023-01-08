package routers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/controllers"
	"go-api-publish/app/repositories"
	"go-api-publish/app/services"
)

var (
	CategoryRepository repositories.CategoryRepositoryInterface = repositories.NewCategory()
	CategoryService    services.CategoryServiceInterface        = services.NewCategory(CategoryRepository)
	CategoryController controllers.CategoryControllerInterface  = controllers.NewCategory(CategoryService)

	SliderRepository repositories.SliderRepositoryInterface = repositories.NewSlider()
	SliderService    services.SliderServiceInterface        = services.NewSlider(SliderRepository)
	SliderController controllers.SliderControllerInterface  = controllers.NewSlider(SliderService)
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

}

package routers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/controllers"
	"go-api-publish/app/repositories"
	"go-api-publish/app/services"
)

var (
	CategoryRepository repositories.CategoryRepositoryInterface = repositories.New()
	CategoryService    services.CategoryServiceInterface        = services.New(CategoryRepository)
	CategoryController controllers.CategoryControllerInterface  = controllers.New(CategoryService)
)

func ApiRouter(group *gin.RouterGroup) {
	cateGroup := group.Group("category")
	{
		cateGroup.GET("/",CategoryController.GetAll)

	}

}

package services

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/helpers"
	"go-api-publish/app/models"
	"go-api-publish/app/repositories"
	"net/http"
)

type CategoryServiceInterface interface {
	GetAll(ctx *gin.Context)
}
type CategoryService struct {
	categoryRepo repositories.CategoryRepositoryInterface
}

func NewCategory(categoryRepo repositories.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}
func (categoryService *CategoryService) GetAll(ctx *gin.Context) {
	var category []models.Category
	err := categoryService.categoryRepo.GetAll(ctx, &category)
	if err != nil {
		res := helpers.ApiResponseError(nil, err, 401)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ApiResponse(category, "List category", 200)
	ctx.JSON(http.StatusOK, res)

}

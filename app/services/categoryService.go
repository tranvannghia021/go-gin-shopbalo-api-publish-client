package services

import (
	"github.com/gin-gonic/gin"
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

func New(categoryRepo repositories.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}
func (categoryService *CategoryService) GetAll(ctx *gin.Context) {
	var category []models.Category
	//var w http.ResponseWriter
	err := categoryService.categoryRepo.GetAll(ctx, &category)
	if err != nil {
		//helpers.ApiResponseError(w, http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	//category =ctx.ShouldBindJSON()
	//helpers.ApiResponse(w, http.StatusOK, category)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    category,
		"message": "List category",
	})
}

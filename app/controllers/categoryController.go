package controllers

import (
	"github.com/gin-gonic/gin"
	_ "go-api-publish/app/models"
	"go-api-publish/app/services"
)

type CategoryControllerInterface interface {
	GetAll(ctx *gin.Context)
}
type CategoryController struct {
	CategoryService services.CategoryServiceInterface
}

func New(CategoryService services.CategoryServiceInterface) CategoryControllerInterface {
	return &CategoryController{
		CategoryService: CategoryService,
	}
}
func (CategoryController *CategoryController) GetAll(ctx *gin.Context) {
	CategoryController.CategoryService.GetAll(ctx)

}

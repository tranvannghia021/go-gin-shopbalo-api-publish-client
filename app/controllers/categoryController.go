package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/services"
)

type CategoryControllerInterface interface {
	GetAll(ctx *gin.Context)
}
type CategoryController struct {
	CategoryServiceInterface services.CategoryServiceInterface
}

func NewCategory(CategoryServiceInterface services.CategoryServiceInterface) CategoryControllerInterface {
	return &CategoryController{
		CategoryServiceInterface: CategoryServiceInterface,
	}
}
func (CategoryController *CategoryController) GetAll(ctx *gin.Context) {

	CategoryController.CategoryServiceInterface.GetAll(ctx)

}

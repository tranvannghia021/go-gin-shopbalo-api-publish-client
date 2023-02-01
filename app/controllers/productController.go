package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-publish/app/services"
)

type ProductControllerInterface interface {
	GetAll(ctx *gin.Context)
}
type ProductController struct {
	productServiceInterface services.ProductServiceInterface
}

func NewProduct(productServiceInterface services.ProductServiceInterface) ProductControllerInterface {
	return &ProductController{
		productServiceInterface: productServiceInterface,
	}
}
func (p *ProductController) GetAll(ctx *gin.Context) {
	p.productServiceInterface.GetAll(ctx)
}

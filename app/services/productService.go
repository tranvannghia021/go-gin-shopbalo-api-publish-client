package services

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-publish/app/helpers"
	"go-api-publish/app/models"
	"go-api-publish/app/repositories"
	"net/http"
	"os"
	"strconv"
)

type ProductServiceInterface interface {
	GetAll(ctx *gin.Context)
}
type ProductService struct {
	productRepositoryInterface repositories.ProductRepositoryInterface
}

func NewProduct(productRepositoryInterface repositories.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		productRepositoryInterface: productRepositoryInterface,
	}
}

type RatingResponse struct {
	Point int `json:"point"`
}
type ProductResponse struct {
	ID            uint32 `json:"id"`
	Category_id   uint32 `json:"category_id"`
	Category_name string `json:"category_name"`
	Name          string `json:"name"`
	Status        int    `json:"status"`
	Description   string `json:"description"`
	Image         string `json:"image"`
	Image_slide   string `json:"image_slide"`
	Code_color    string `json:"code_color"`
	Amount        int    `json:"amount"`
	Price         int    `json:"price"`
	Ratings       RatingResponse
}

func (p *ProductService) GetAll(ctx *gin.Context) {
	limit := 10
	currentPage := 1
	perPage, _ := strconv.Atoi(ctx.Query("per_page"))
	sort := ctx.Query("sort")

	page, _ := strconv.Atoi(ctx.Query("page"))
	if perPage == 0 {
		perPage = limit
	}
	if page == 0 {
		page = currentPage
	}
	if sort == "" {
		sort = "asc"
	}
	godotenv.Load()
	host := os.Getenv("APP_URL") + ctx.FullPath()
	//ctx.JSON(http.StatusBadRequest, gin.H{
	//	"data": perPage,
	//})
	//return
	var product []models.Product
	err := p.productRepositoryInterface.GetAll(ctx, &product, sort)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ApiResponseError(nil, err, 401))
		return
	}
	var ProductRes []interface{}

	for _, item := range product {
		if item.ProductDetail.Product_id < 0 {
			continue
		}
		var avg float64 = 0

		lenPoint := len(item.Rating)
		if lenPoint > 0 {
			for _, rating := range item.Rating {
				avg += float64(rating.Point)
			}
			avg = (avg) / float64(lenPoint)
		}
		if avg == 0 {
			avg = 5
		}

		ProductRes = append(ProductRes, ProductResponse{
			ID:            item.ID,
			Category_id:   item.Category_id,
			Category_name: item.Category.Name,
			Name:          item.Name,
			Status:        item.Status,
			Description:   item.Description,
			Image:         item.Image,
			Image_slide:   item.Image_slide,
			Code_color:    item.ProductDetail.Code_color,
			Amount:        item.ProductDetail.Amount,
			Price:         item.ProductDetail.Price,
			Ratings: RatingResponse{
				Point: int(avg),
			},
		})
	}
	res := helpers.Paginate(ProductRes, perPage, page, host)

	ctx.JSON(http.StatusOK, helpers.ApiResponsePaginate(res, "list product", 200))
	return
	//ctx.JSON(http.StatusOK, helpers.ApiResponse(pg, "list product", 200))
}

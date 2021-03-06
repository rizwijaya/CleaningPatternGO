package product

import (
	"ClearningPatternGO/modules/v1/utilities/product/repository"
	"ClearningPatternGO/modules/v1/utilities/product/service"
	api "ClearningPatternGO/pkg/api_response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler interface {
	ListProduct(c *gin.Context)
}

type productHandler struct {
	productService service.Service
}

func NewProductHandler(productService service.Service) *productHandler {
	return &productHandler{productService}
}

func Handler(db *gorm.DB) *productHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := NewProductHandler(Service)
	return Handler
}

func (h *productHandler) ListProduct(c *gin.Context) {
	listProduct, err := h.productService.ListProduct()

	if err != nil {
		response := api.APIRespon("Failed to get data all product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := api.APIRespon("Success to get data all product", http.StatusOK, "success", listProduct)
	c.JSON(http.StatusOK, response)
}

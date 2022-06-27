package product

import (
	"ClearningPatternGO/app/helper"
	"ClearningPatternGO/modules/v1/utilities/product/repository"
	"ClearningPatternGO/modules/v1/utilities/product/service"
	ss "ClearningPatternGO/modules/v1/utilities/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler interface {
	ListProduct(c *gin.Context)
}

type productHandler struct {
	productService ss.Service
}

func NewProductHandler(productService ss.Service) *productHandler {
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
		response := helper.APIRespon("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespon("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

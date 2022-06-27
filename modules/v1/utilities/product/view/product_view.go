package view

import (
	"ClearningPatternGO/modules/v1/utilities/product/repository"
	"ClearningPatternGO/modules/v1/utilities/product/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productView struct {
	productService service.Service
}

func NewProductView(productService service.Service) *productView {
	return &productView{productService}
}

func View(db *gorm.DB) *productView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	View := NewProductView(Service)
	return View
}

func (h *productView) Index(c *gin.Context) {
	device, err := h.productService.GetAllProduct()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "device_index.html", gin.H{"device": device})
}

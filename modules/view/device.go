package view

import (
	"ClearningPatternGO/modules/utilities/device"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deviceView struct {
	deviceService device.Service
}

func NewDeviceView(deviceService device.Service) *deviceView {
	return &deviceView{deviceService}
}

func (h *deviceView) Index(c *gin.Context) {
	device, err := h.deviceService.GetAllDevice()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "device_index.html", gin.H{"device": device})
}

func (h *deviceView) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	device, err := h.deviceService.GetDeviceByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "device_edit.html", gin.H{"device": device})
}

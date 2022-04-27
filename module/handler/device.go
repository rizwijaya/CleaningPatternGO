package handler

import (
	"TamaskaPJU/module/utilities/device"

	"github.com/gin-gonic/gin"
)

type DeviceHandler interface {
	ViewDataDevice(c *gin.Context)
}

type deviceHandler struct {
	deviceService device.Service
}

func NewDeviceHandler(deviceService device.Service) *deviceHandler {
	return &deviceHandler{deviceService}
}

func (h *deviceHandler) ViewDataDevice(c *gin.Context) {
	h.deviceService.GetAllDevice()
}

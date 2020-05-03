package controller

import (
	"github.com/devingen/atama-api/controller/service-controller"
	"github.com/devingen/atama-api/service"
)

// NewServiceController generates new ServiceController
func NewServiceController(service service.AtamaService) *service_controller.ServiceController {
	return &service_controller.ServiceController{
		Service: service,
	}
}

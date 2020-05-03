package service_controller

import (
	"github.com/devingen/atama-api/service"
)

// ServiceController implements AtamaController interface by using AtamaService
type ServiceController struct {
	Service service.AtamaService
}

package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos/requests"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type AlertHandler struct {
	alertService services.IAlertService
	service      *apiservicemanager.ApiService
}

func NewAlertHandler(apiService *apiservicemanager.ApiService, alertService services.IAlertService) *AlertHandler {
	return &AlertHandler{
		alertService: alertService,
		service:      apiService,
	}
}

func (h *AlertHandler) SetupRoutes() {
	h.service.GET("", h.ListAlerts)
	h.service.GET("/:id", h.GetAlertById)
}

func (h *AlertHandler) ListAlerts(c echo.Context) error {
	alerts, err := h.alertService.GetAllAlerts()

	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, alerts)
}

func (h *AlertHandler) GetAlertById(c echo.Context) error {
	var req requests.IdParam

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Missing Id"))
	}

	alert, err := h.alertService.GetAlertById(req.Id)
	if err != nil {
		return c.JSON(404, dtos.ErrorResponseMsg("Alert not found"))
	}

	return c.JSON(200, alert)
}

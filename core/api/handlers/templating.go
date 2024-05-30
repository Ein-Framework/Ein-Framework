package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type TemplatingHandler struct {
	templatingManager templating.ITemplateManager
}

func New(templatingManager templating.ITemplateManager) *TemplatingHandler {
	return &TemplatingHandler{
		templatingManager: templatingManager,
	}
}

func (h *TemplatingHandler) SetupRoutes(service apiservicemanager.ApiService) {
	service.GET("/", h.GetAllAvailableTemplates)
	service.GET("/loaded", h.GetAllLoadedTemplates)
}

func (h *TemplatingHandler) GetAllAvailableTemplates(c echo.Context) error {
	templates, err := h.templatingManager.ListAllAvailableTemplates()
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(templates))
}

func (h *TemplatingHandler) GetAllLoadedTemplates(c echo.Context) error {
	templates := h.templatingManager.GetAllLoadedTemplatesMeta()
	return c.JSON(200, dtos.SuccessDataMsgResponse(templates))
}

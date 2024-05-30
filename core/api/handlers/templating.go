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

func NewTemplatingHandler(templatingManager templating.ITemplateManager) *TemplatingHandler {
	return &TemplatingHandler{
		templatingManager: templatingManager,
	}
}

func (h *TemplatingHandler) SetupTemplatingRoutes(service *apiservicemanager.ApiService) {
	service.GET("/", h.GetAllAvailableTemplates)
	service.GET("/loaded", h.GetAllLoadedTemplates)

	service.POST("/load", h.LoadTemplate)
	service.DELETE("/unload", h.UnloadTemplate)
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

func (h *TemplatingHandler) LoadTemplate(c echo.Context) error {
	req := &dtos.LoadTemplateRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	template, err := h.templatingManager.LoadTemplate(req.Name)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(
		dtos.LoadTemplateResponse{
			Name: req.Name,
			Meta: template.Meta,
		},
	))
}

func (h *TemplatingHandler) UnloadTemplate(c echo.Context) error {
	req := &dtos.UnloadTemplateRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	err := h.templatingManager.UnloadTemplate(req.Name)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(
		dtos.UnloadTemplateResponse{
			Name: req.Name,
		},
	))
}

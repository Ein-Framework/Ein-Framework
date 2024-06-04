package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type TemplatingHandler struct {
	service           *apiservicemanager.ApiService
	templatingManager templating.ITemplateManager
}

func NewTemplatingHandler(service *apiservicemanager.ApiService, templatingManager templating.ITemplateManager) *TemplatingHandler {
	return &TemplatingHandler{
		service:           service,
		templatingManager: templatingManager,
	}
}

func (h *TemplatingHandler) SetupRoutes() {
	h.service.GET("/", h.GetAllAvailableTemplates)
	h.service.GET("/loaded", h.GetAllLoadedTemplates)

	h.service.POST("/load", h.LoadTemplate)
	h.service.DELETE("/unload", h.UnloadTemplate)

	h.service.GET("/job/:id", h.GetJobTemplates)
}

func (h *TemplatingHandler) GetJobTemplates(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	templates, err := h.templatingManager.GetTemplatesForJob(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(templates))
}

func (h *TemplatingHandler) GetAllAvailableTemplates(c echo.Context) error {
	templates, err := h.templatingManager.GetAllAvailableTemplates()
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

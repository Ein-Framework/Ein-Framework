package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/Ein-Framework/Ein-Framework/pkg/plugins"
	"github.com/labstack/echo/v4"
)

type PluginsHandler struct {
	pluginManager plugins.IPluginManager
}

func NewPluginsHandler(pluginManager plugins.IPluginManager) *PluginsHandler {
	return &PluginsHandler{
		pluginManager: pluginManager,
	}
}

func (h *PluginsHandler) SetupPluginRoutes(service *apiservicemanager.ApiService) {
	service.GET("/", h.GetAllAvailablePlugins)
	service.GET("/loaded", h.GetAllLoadedPlugins)

	service.POST("/load", h.LoadPlugin)
	service.DELETE("/unload", h.UnloadPlugin)
}

func (h *PluginsHandler) GetAllAvailablePlugins(c echo.Context) error {
	plugins, err := h.pluginManager.ListAllPlugins()
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(plugins))
}

func (h *PluginsHandler) GetAllLoadedPlugins(c echo.Context) error {
	templates := h.pluginManager.ListLoadedPlugins()
	return c.JSON(200, dtos.SuccessDataMsgResponse(templates))
}

func (h *PluginsHandler) LoadPlugin(c echo.Context) error {
	req := &dtos.LoadPluginRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	loadedPlugin, err := h.pluginManager.LoadPlugin(req.Name)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}
	return c.JSON(200, dtos.SuccessDataMsgResponse(
		dtos.LoadPluginResponse{
			Name: req.Name,
			Meta: loadedPlugin.Plugin.MetaInfo(),
			Info: loadedPlugin.Plugin.Info(),
		},
	))
}

func (h *PluginsHandler) UnloadPlugin(c echo.Context) error {
	req := &dtos.UnloadPluginRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	err := h.pluginManager.UnloadPlugin(req.Name)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(
		dtos.UnloadPluginResponse{
			Name: req.Name,
		},
	))
}

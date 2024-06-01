package api

import (
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	"github.com/labstack/echo/v4"
)

type AppComponents struct {
	TemplatingManager templating.ITemplateManager
}

type ApiServer struct {
	server     *echo.Echo
	components *AppComponents
}

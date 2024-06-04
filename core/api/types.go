package api

import (
	taskmanager "github.com/Ein-Framework/Ein-Framework/core/task_manager"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
	"github.com/labstack/echo/v4"
)

type AppComponents struct {
	TemplatingManager templating.ITemplateManager
	TaskManager       taskmanager.ITaskManager
}

type ApiServer struct {
	server     *echo.Echo
	components *AppComponents
}

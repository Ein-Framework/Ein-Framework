package api

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/api/handlers"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(coreServices *services.Services, components *AppComponents, config *config.Config, logger *zap.Logger) *ApiServer {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// assessmentHandler := handlers.NewAssessmentHandler(assessmentService)

	serviceManager := apiservicemanager.NewServiceManager(e)

	templatingService, err := serviceManager.NewService("templating")
	if err != nil {
		logger.Panic("failed to create templating error")
	}

	pluginService, err := serviceManager.NewService("plugin")
	if err != nil {
		logger.Panic("failed to create templating error")
	}

	templatingHandler := handlers.NewTemplatingHandler(components.TemplatingManager)
	pluginHandler := handlers.NewPluginsHandler(components.TemplatingManager.PluginManager())

	templatingHandler.SetupTemplatingRoutes(templatingService)
	pluginHandler.SetupPluginRoutes(pluginService)

	// e.PUT("/assessments/:id", assessmentHandler.UpdateAssessment)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerHTTPPort)))

	return &ApiServer{
		server:     e,
		components: components,
	}
}

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
	err := coreServices.AssessmentStageService.InitStages()
	if err != nil {
		logger.Panic("failed to init stages")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	serviceManager := apiservicemanager.NewServiceManager(e)

	assessmentService, err := serviceManager.NewService(("assessment"))
	if err != nil {
		logger.Panic("failed to create assessment service")
	}

	templatingService, err := serviceManager.NewService("templating")
	if err != nil {
		logger.Panic("failed to create templating api service error")
	}

	pluginService, err := serviceManager.NewService("plugin")
	if err != nil {
		logger.Panic("failed to create plugin api service error")
	}

	jobExecutionService, err := serviceManager.NewService("execution")
	if err != nil {
		logger.Panic("failed to create job execution api service error")
	}

	jobService, err := serviceManager.NewService("job")
	if err != nil {
		logger.Panic("failed to create job api service error")
	}

	alertService, err := serviceManager.NewService("alert")
	if err != nil {
		logger.Panic("failed to create alert api service error")
	}

	templatingHandler := handlers.NewTemplatingHandler(templatingService, components.TemplatingManager)
	pluginHandler := handlers.NewPluginsHandler(pluginService, components.TemplatingManager.PluginManager())
	jobExecutionHandler := handlers.NewJobExecutionHandler(jobExecutionService, components.TaskManager)
	jobHandler := handlers.NewJobHandler(jobService, coreServices.JobService)
	assessmentHandler := handlers.NewAssessmentHandler(assessmentService, coreServices.AssessmentService)
	alertHandler := handlers.NewAlertHandler(alertService, coreServices.AlertService)

	SetupRoutes(
		templatingHandler,
		pluginHandler,
		jobExecutionHandler,
		jobHandler,
		assessmentHandler,
		alertHandler,
	)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerHTTPPort)))

	return &ApiServer{
		server:     e,
		components: components,
	}
}

func SetupRoutes(handlers ...handlers.IHandler) {
	for _, handler := range handlers {
		handler.SetupRoutes()
	}
}

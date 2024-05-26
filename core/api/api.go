package api

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(coreServices *services.Services, components *AppComponents, config *config.Config, logger *zap.Logger) *ApiService {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// assessmentHandler := handlers.NewAssessmentHandler(assessmentService)

	// e.PUT("/assessments/:id", assessmentHandler.UpdateAssessment)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerHTTPPort)))

	return &ApiService{
		server:     e,
		components: components,
	}
}

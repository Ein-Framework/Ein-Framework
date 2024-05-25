package api

import (
	"fmt"
	"log"

	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiService struct {
}

func New(config *config.Config, logger *zap.Logger) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := domain.NewDatabase(config.Database)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(db)

	// Initialize services
	// coreServices := services.InitServices(db, logger, config)

	// assessmentHandler := handlers.NewAssessmentHandler(assessmentService)

	// e.PUT("/assessments/:id", assessmentHandler.UpdateAssessment)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerHTTPPort)))

}

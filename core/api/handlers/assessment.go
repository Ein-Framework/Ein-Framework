package handlers

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos/requests"
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type AssessmentHandler struct {
	assessmentService services.IAssessmentService
	service           *apiservicemanager.ApiService
}

func NewAssessmentHandler(apiService *apiservicemanager.ApiService, assessmentService services.IAssessmentService) *AssessmentHandler {
	return &AssessmentHandler{
		assessmentService: assessmentService,
		service:           apiService,
	}
}

func (h *AssessmentHandler) SetupRoutes() {
	h.service.GET("", h.ListAssesments)
	h.service.GET("/:id", h.GetAssessmentById)
	h.service.POST("", h.CreateAssessment)
	h.service.POST("/url", h.AddNewAssessmentFromURL)
	h.service.PUT("/:id", h.UpdateAssessment)
	h.service.DELETE("/:id", h.DeleteAssessment)
}

func (h *AssessmentHandler) ListAssesments(c echo.Context) error {

	assessments, err := h.assessmentService.GetAllAssessments()
	fmt.Println(assessments)

	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(assessments))
}

func (h *AssessmentHandler) GetAssessmentById(c echo.Context) error {

	var req requests.IdParam

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Missing Id"))
	}

	assessment, err := h.assessmentService.GetAssessmentById(req.Id)
	if err != nil {
		return c.JSON(404, dtos.ErrorResponseMsg("Assessment not found"))
	}

	return c.JSON(200, assessment)
}

func (h *AssessmentHandler) CreateAssessment(c echo.Context) error {
	req := &requests.AssessmentRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	assessment, err := h.assessmentService.AddNewAssessment(
		req.Name,
		req.AssessmentType,
		req.Scope,
		req.EngagementRules,
	)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(assessment))
}

func (h *AssessmentHandler) DeleteAssessment(c echo.Context) error {
	req := &requests.IdParam{}
	var err error

	if err = c.Bind(req); err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	assessment, err := h.assessmentService.GetAssessmentById(req.Id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	err = h.assessmentService.DeleteAssessment(req.Id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(assessment))
}

func (h *AssessmentHandler) UpdateAssessment(c echo.Context) error {

	var err error
	req := &requests.AssessmentRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	updatedAssessment := &entity.Assessment{
		Name:  req.Name,
		Type:  req.AssessmentType,
		Scope: req.Scope,
		// ... other fields
	}

	if err = h.assessmentService.UpdateAssessment(req.Id, updatedAssessment); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, updatedAssessment)
}

func (h *AssessmentHandler) AddNewAssessmentFromURL(c echo.Context) error {

	req := requests.CreateAssessmentFromURL{}
	var err error
	if err := c.Bind(&req); err != nil {
		fmt.Print(err.Error())
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Request"))
	}

	var assessment *entity.Assessment

	switch req.Platform {
	case "hackerone":
		assessment, err = h.assessmentService.AddNewAssessmentFromHackerone(req.ProgramName)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(500, dtos.ErrorResponseMsg("Error creating assessment"))
		}
		return c.JSON(201, assessment)

	default:
		return c.JSON(400, dtos.ErrorResponseMsg("Unsupported Platform"))
	}
}

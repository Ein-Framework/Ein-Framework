package handlers

import (
	"strconv"

	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/labstack/echo/v4"
)

type AssessmentHandler struct {
	assessmentService *services.AssessmentService
}

type AssessmentRequest struct {
	Name           string                `json:"name"`
	AssessmentType entity.AssessmentType `json:"assessmentType"`
	Scope          entity.Scope          `json:"scope"`
}

func NewAssessmentHandler(assessmentService *services.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{
		assessmentService: assessmentService,
	}
}

func (h *AssessmentHandler) CreateAssessment(c echo.Context) error {
	req := &AssessmentRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	_, err := h.assessmentService.AddNewAssessment(req.Name, req.AssessmentType, req.Scope)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Assessment created successfully",
	})
}

func (h *AssessmentHandler) DeleteAssessment(c echo.Context) error {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))

	}
	err = h.assessmentService.DeleteAssessment(uint(uid))
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.InfoMsgResponse("Assessment deleted successfully"))
}

func (h *AssessmentHandler) UpdateAssessment(c echo.Context) error {
	id := c.Param("id")

	req := &AssessmentRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	updatedAssessment := &entity.Assessment{
		Name:  req.Name,
		Type:  req.AssessmentType,
		Scope: req.Scope,
		// ... other fields
	}
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))

	}

	err = h.assessmentService.UpdateAssessment(uint(uid), updatedAssessment)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.InfoMsgResponse("Assessment updated successfully"))
}

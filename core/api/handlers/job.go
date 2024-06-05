package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type JobHandler struct {
	service    *apiservicemanager.ApiService
	jobService services.IJobService
}

func NewJobHandler(service *apiservicemanager.ApiService, jobService services.IJobService) *JobHandler {
	return &JobHandler{
		service:    service,
		jobService: jobService,
	}
}

func (h *JobHandler) SetupRoutes() {
	h.service.GET("/:id", h.GetJobById)
	h.service.GET("", h.GetAllJobs)

	h.service.POST("", h.CreateJob)

	h.service.PATCH("/:id", h.UpdateJob)

	h.service.DELETE("/:id", h.DeleteJob)
}

func (h *JobHandler) GetJobById(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	job, err := h.jobService.GetJobById(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(job))
}

func (h *JobHandler) GetAllJobs(c echo.Context) error {
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(jobs))
}

func (h *JobHandler) CreateJob(c echo.Context) error {
	req := &dtos.CreateJobRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	job, err := h.jobService.AddNewJob(req.Name, req.Templates)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(job))
}

func (h *JobHandler) UpdateJob(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	req := &dtos.UpdateJobRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	err = h.jobService.UpdateJob(id, &entity.Job{
		Name: req.Name,
	})
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	h.jobService.UpdateJobTemplates(id, req.Templates...)

	newJob, err := h.jobService.GetJobById(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(newJob))
}

func (h *JobHandler) DeleteJob(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	oldJob, err := h.jobService.GetJobById(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	err = h.jobService.DeleteJob(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(oldJob))
}

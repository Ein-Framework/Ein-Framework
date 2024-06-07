package handlers

import (
	"github.com/Ein-Framework/Ein-Framework/core/api/dtos"
	taskmanager "github.com/Ein-Framework/Ein-Framework/core/task_manager"
	apiservicemanager "github.com/Ein-Framework/Ein-Framework/pkg/api_service_manager"
	"github.com/labstack/echo/v4"
)

type JobExecutionHandler struct {
	service     *apiservicemanager.ApiService
	taskManager taskmanager.ITaskManager
}

func NewJobExecutionHandler(service *apiservicemanager.ApiService, taskManager taskmanager.ITaskManager) *JobExecutionHandler {
	return &JobExecutionHandler{
		service:     service,
		taskManager: taskManager,
	}
}

func (h *JobExecutionHandler) SetupRoutes() {
	h.service.GET("/job/:id/state", h.GetJobState)
	h.service.GET("/task/:id/state", h.GetTaskState)
	h.service.GET("/active", h.GetRunningJobs)

	h.service.POST("/job", h.ExecuteJob)
	h.service.POST("/template", h.ExecuteTemplate)
	h.service.GET("/active", h.GetRunningJobs)

	h.service.DELETE("/:id", h.CancelJobExecution)

	h.service.GET("", h.GetAllJobs)
}

func (h *JobExecutionHandler) GetJobState(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	state, err := h.taskManager.ViewJobStatus(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(*state))
}

func (h *JobExecutionHandler) GetTaskState(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	state, err := h.taskManager.ViewTaskStatus(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(*state))
}

func (h *JobExecutionHandler) ExecuteJob(c echo.Context) error {
	req := &dtos.RunJobRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	exec, err := h.taskManager.ExecuteJob(req.JobId, req.AssessmentId)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(exec))
}

func (h *JobExecutionHandler) ExecuteTemplate(c echo.Context) error {
	req := &dtos.RunTemplateRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	task, err := h.taskManager.ExecuteTemplate(req.Template, req.AssessmentId)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(task))
}

func (h *JobExecutionHandler) CancelJobExecution(c echo.Context) error {
	id, err := GetUIntParam(c, "id")
	if err != nil {
		return c.JSON(400, dtos.ErrorResponseMsg("Bad Id"))
	}

	err = h.taskManager.CancelJob(id)
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.InfoMsgResponse("Job execution canceled"))
}

func (h *JobExecutionHandler) GetAllJobs(c echo.Context) error {
	executions, err := h.taskManager.GetAllJobs()
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(executions))
}

func (h *JobExecutionHandler) GetRunningJobs(c echo.Context) error {
	executions, err := h.taskManager.GetRunningJobs()
	if err != nil {
		return c.JSON(500, dtos.ErrorResponseMsg(err.Error()))
	}

	return c.JSON(200, dtos.SuccessDataMsgResponse(executions))
}

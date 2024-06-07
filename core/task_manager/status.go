package taskmanager

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
)

func (manager *TaskManager) ViewJobStatus(jobID uint) (*entity.TaskState, error) {
	fmt.Println("ViewJobStatus")
	jobExec, err := manager.coreServices.JobExecutionService.GetJobExecutionById(jobID)
	if err != nil {
		return nil, err
	}

	return &jobExec.Status, nil
}

func (manager *TaskManager) ViewTaskStatus(taskID uint) (*entity.TaskState, error) {
	task, err := manager.coreServices.TaskService.GetTaskById(taskID)
	if err != nil {
		return nil, err
	}

	return &task.State, nil
}

func (manager *TaskManager) GetRunningJobs() ([]*entity.JobExecution, error) {
	result := make([]*entity.JobExecution, 0)
	manager.executionsMutex.Lock()
	defer manager.executionsMutex.Unlock()

	for _, v := range manager.executions {
		result = append(result, v.jobExecution)
	}

	return result, nil
}

func (manager *TaskManager) GetAllJobs() ([]*entity.JobExecution, error) {
	res, err := manager.coreServices.JobExecutionService.GetAllJobExecutions()
	if err != nil {
		return nil, err
	}

	return res, nil
}

package taskmanager

import (
	"time"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

type ConcurrentExecution struct {
	jobExecution *entity.JobExecution // Optional, can be nil for single template execution
	tasks        queue.FifoQueue[entity.Task]
}

func (manager *TaskManager) execute(exec *ConcurrentExecution) {
	for {
		task, err := exec.tasks.Remove()
		if err != nil {
			break
		}

		manager.coreServices.TaskService.UpdateTaskState(task.ID, entity.Running)

		for _, asset := range task.Assessment.Assets {
			manager.templateManager.ExecuteTemplate(string(task.Template), CreateExecutionContext(*task, asset))
		}

		manager.coreServices.TaskService.UpdateTaskState(task.ID, entity.Stopped)

		if exec.jobExecution != nil {
			time.Sleep(time.Second * time.Duration(exec.jobExecution.Assessment.EngagementRules.RateLimitPerSecond))
		}
	}
}

func (manager *TaskManager) RunConcurrentExecution(exec *ConcurrentExecution) {
	go func() {
		if exec.jobExecution != nil {
			jobExec := exec.jobExecution

			manager.AddConcurrentExecution(jobExec.ID, exec)

			manager.coreServices.JobExecutionService.UpdateJobExecutionStatus(jobExec.ID, entity.Running)
			defer func() {
				manager.RemoveConcurrentExecution(exec.jobExecution.ID)

				manager.coreServices.JobExecutionService.UpdateJobExecutionStatus(jobExec.ID, entity.Stopped)
			}()
		}
		manager.execute(exec)
	}()
}

func (manager *TaskManager) AddConcurrentExecution(execId uint, exec *ConcurrentExecution) {
	manager.executionsMutex.Lock()
	defer manager.executionsMutex.Unlock()
	manager.executions[execId] = exec
}

func (manager *TaskManager) RemoveConcurrentExecution(execId uint) {
	manager.executionsMutex.Lock()
	defer manager.executionsMutex.Unlock()
	delete(manager.executions, execId)
}

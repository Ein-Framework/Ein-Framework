package taskmanager

import (
	"time"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

type ConcurrentExecution struct {
	jobExecution *entity.JobExecution
	tasks        queue.FifoQueue[entity.Task]
}

func (exec *ConcurrentExecution) execute() {
	for {
		_, err := exec.tasks.Remove()
		if err == nil {
			break
		}

		if exec.jobExecution != nil {
			time.Sleep(time.Second * time.Duration(exec.jobExecution.Assessment.EngagementRules.RateLimitPerSecond))
		}
	}
}

func (manager *TaskManager) RunConcurrentExecution(exec *ConcurrentExecution) {
	go func() {
		if exec.jobExecution != nil {
			manager.AddConcurrentExecution(exec.jobExecution.ID, exec)
			defer func() {
				manager.RemoveConcurrentExecution(exec.jobExecution.ID)
			}()
		}
		exec.execute()
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

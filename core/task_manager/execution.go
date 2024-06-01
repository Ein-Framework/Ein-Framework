package taskmanager

import (
	"errors"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

func (manager *TaskManager) ExecuteTemplate(Template entity.Template, AssesementID uint) (*entity.Task, error) {
	_, err := manager.coreServices.AssessmentService.GetAssessmentById(AssesementID)
	if err != nil {
		return nil, err
	}
	return manager.createTask(Template, AssesementID)
}

func (manager *TaskManager) createTask(Template entity.Template, AssesementID uint) (*entity.Task, error) {
	task, err := manager.coreServices.TaskService.AddNewTask(entity.Queued, Template, AssesementID)
	if err != nil {
		return nil, err
	}
	exec := &ConcurrentExecution{
		tasks: queue.CreateQueue[entity.Task](),
	}
	err = exec.tasks.Insert(task)
	if err != nil {
		return nil, err
	}
	manager.RunConcurrentExecution(exec)
	return task, nil
}

func (manager *TaskManager) ExecuteJob(jobID uint, AssesementID uint) (*entity.JobExecution, error) {
	job, err := manager.coreServices.JobService.GetJobById(jobID)
	if err != nil {
		return nil, err
	}
	_, err = manager.coreServices.AssessmentService.GetAssessmentById(AssesementID)
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0)
	tasksQueue := queue.CreateQueue[entity.Task]()
	for _, template := range job.Templates {
		task, err := manager.createTask(template, AssesementID)
		if err != nil {
			manager.coreServices.TaskService.DeleteTasks(tasks...)
			return nil, err
		}
		tasks = append(tasks, *task)
		tasksQueue.Insert(task)
	}

	jobExec, err := manager.coreServices.JobExecutionService.AddNewJobExecution(jobID, AssesementID, tasks, entity.Queued)
	if err != nil {
		manager.coreServices.TaskService.DeleteTasks(tasks...)
		return nil, err
	}

	exec := &ConcurrentExecution{
		tasks:        tasksQueue,
		jobExecution: jobExec,
	}
	manager.RunConcurrentExecution(exec)
	return jobExec, nil
}

func (manager *TaskManager) CancelJob(jobID uint) error {
	jobExec, err := manager.coreServices.JobExecutionService.GetJobExecutionById(jobID)
	if err != nil {
		return err
	}

	manager.executionsMutex.Lock()
	conc, ok := manager.executions[jobID]
	if !ok {
		return errors.New("job execution already ended")
	}
	tasks := conc.tasks.Empty()
	manager.executionsMutex.Unlock()

	jobExec.Status = entity.Canceled
	manager.coreServices.JobExecutionService.UpdateJobExecution(jobExec.ID, jobExec)

	for _, task := range tasks {
		task.State = entity.Canceled
		manager.coreServices.TaskService.UpdateTask(task.ID, task)
	}
	return nil
}

func (manager *TaskManager) ViewJobStatus(jobID uint) (*entity.TaskState, error) {
	jobExec, err := manager.coreServices.JobExecutionService.GetJobExecutionById(jobID)
	if err != nil {
		return nil, err
	}

	return &jobExec.Status, nil
}

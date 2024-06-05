package taskmanager

import (
	"errors"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/queue"
)

func (manager *TaskManager) ExecuteTemplate(Template entity.Template, AssesementID uint) (*entity.Task, error) {
	assessment, err := manager.coreServices.AssessmentService.GetAssessmentById(AssesementID)
	if err != nil {
		return nil, err
	}
	return manager.createTask(Template, AssesementID, assessment.StageID)
}

func (manager *TaskManager) createTask(Template entity.Template, AssesementID uint, assessmentStageId uint) (*entity.Task, error) {
	task, err := manager.coreServices.TaskService.AddNewTask(entity.Queued, Template, AssesementID, assessmentStageId)
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

	if len(job.Templates) == 0 {
		return nil, errors.New("no templates associated to this job")
	}

	assessment, err := manager.coreServices.AssessmentService.GetAssessmentById(AssesementID)
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0)
	tasksQueue := queue.CreateQueue[entity.Task]()
	for _, template := range job.Templates {
		task := entity.Task{
			Template:          template.Template,
			AssessmentId:      AssesementID,
			AssessmentStageId: assessment.StageID,
			State:             entity.Queued,
		}

		tasks = append(tasks, task)
		tasksQueue.Insert(&task)
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

	manager.coreServices.JobExecutionService.UpdateJobExecutionStatus(jobExec.ID, entity.Canceled)

	for _, task := range tasks {
		manager.coreServices.TaskService.UpdateTaskState(task.ID, entity.Canceled)
	}
	return nil
}

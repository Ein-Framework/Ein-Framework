package taskmanager

import (
	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
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
	err = manager.q.Insert(task)
	if err != nil {
		return nil, err
	}
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
	for _, template := range job.Templates {
		task, err := manager.createTask(template, AssesementID)
		if err != nil {
			//TODO : yaaaaaaaaaaa mongi delete tasks
			return nil, err
		}
		tasks = append(tasks, *task)
	}

	jobExec, err := manager.coreServices.JobExecutionService.AddNewJobExecution(jobID, AssesementID, tasks, entity.Queued)
	if err != nil {
		//TODO : yaaaaaaaaaaa mongi delete tasks
		return nil, err
	}
	return jobExec, nil
}

func (manager *TaskManager) CancelJob(jobID uint) error {
	return nil
}

func (manager *TaskManager) ViewJobStatus(jobID uint) (*entity.TaskState, error) {
	return nil, nil
}

package taskmanager

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

func (manager *TaskManager) ExecuteTemplate(Template entity.Template, AssesementID uint) (*entity.Task, error) {
	_, err := manager.coreServices.AssessmentService.GetAssessmentById(AssesementID)
	if err != nil {
		return nil, err
	}

	//create a task
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
	return nil, nil
}

func (manager *TaskManager) CancelJob(jobID uint) error {
	return nil
}

func (manager *TaskManager) ViewJobStatus(jobID uint) (*entity.TaskState, error) {
	return nil, nil
}

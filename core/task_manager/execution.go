package taskmanager

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

func (manager *TaskManager) ExecuteTemplate(TemplateID uint) (*entity.JobExecution, error) {
	return nil, nil
}

func (manager *TaskManager) ExecuteJob(jobID uint) (*entity.JobExecution, error) {
	return nil, nil
}

func (manager *TaskManager) CancelJob(jobID uint) error {
	return nil
}

func (manager *TaskManager) ViewJobStatus(jobID uint) (*entity.TaskState, error) {
	return nil, nil
}

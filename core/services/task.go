package services

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type TaskService struct {
	Service
}

func NewTaskService(ctx Context) *TaskService {
	repo := repository.NewGormRepository(ctx.OrmConnection.Db, ctx.Logger.Sugar())
	return &TaskService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *TaskService) AddNewTask(state entity.TaskState, output string, outputFormat entity.OutputFormat, args map[string]string, assessmentStageId uint) (*entity.Task, error) {
	if state == "" || outputFormat == "" || assessmentStageId == 0 {
		return nil, fmt.Errorf("invalid input: state, outputFormat, and assessmentStageId are required")
	}

	task := &entity.Task{
		State:             state,
		Output:            output,
		OutputFormat:      outputFormat,
		Args:              args,
		AssessmentStageId: assessmentStageId,
	}

	if err := s.repo.Create(task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	task := &entity.Task{Model: gorm.Model{ID: id}}
	if err := s.repo.Delete(task); err != nil {
		return fmt.Errorf("failed to delete task with ID %d: %w", id, err)
	}
	return nil
}

func (s *TaskService) UpdateTask(id uint, updatedTask *entity.Task) error {
	task, err := s.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("task with ID %d not found: %w", id, err)
	}

	task.State = updatedTask.State
	task.Output = updatedTask.Output
	task.OutputFormat = updatedTask.OutputFormat
	task.Args = updatedTask.Args
	task.AssessmentStageId = updatedTask.AssessmentStageId
	task.AssessmentStage = updatedTask.AssessmentStage

	if err := s.repo.Save(task); err != nil {
		return fmt.Errorf("failed to update task with ID %d: %w", id, err)
	}

	return nil
}

func (s *TaskService) GetTaskById(id uint) (*entity.Task, error) {
	var task entity.Task
	if err := s.repo.GetOneByID(&task, id); err != nil {
		return nil, fmt.Errorf("task with ID %d not found: %w", id, err)
	}
	return &task, nil
}

func (s *TaskService) GetAllTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := s.repo.GetAll(&tasks); err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %w", err)
	}
	return tasks, nil
}
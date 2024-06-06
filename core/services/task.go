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
	repo := repository.NewGormRepository(
		ctx.OrmConnection.Db,
		ctx.Logger.Sugar(),
	)
	return &TaskService{
		Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
	}
}

func (s *TaskService) AddNewTask(state entity.TaskState, template entity.Template, assessmentId uint, assessmentStageId uint) (*entity.Task, error) {
	if state == "" || /* outputFormat == "" || */ assessmentStageId == 0 {
		return nil, fmt.Errorf("invalid input: state, outputFormat, and assessmentStageId are required")
	}

	task := &entity.Task{
		State:    state,
		Template: template,
		//Output:            output,
		//OutputFormat:      outputFormat,
		//Args:              args,
		AssessmentId:      assessmentId,
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

func (s *TaskService) DeleteTasks(tasks ...entity.Task) []error {
	errs := make([]error, 0)
	for _, task := range tasks {
		err := s.DeleteTask(task.ID)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
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

func (s *TaskService) UpdateTaskState(id uint, state entity.TaskState) error {
	s.logger.Debug("TaskService: Updating task")

	task, err := s.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("task with ID %d not found: %w", id, err)
	}

	task.State = state
	if err := s.repo.Save(task); err != nil {
		return fmt.Errorf("failed to update task with ID %d: %w", id, err)
	}

	return nil
}

func (s *TaskService) GetTaskById(id uint) (*entity.Task, error) {
	var task entity.Task

	err := s.repo.GetOneByID(
		&task,
		id,
		"Assessment",
		"Assessment.Scope.InScope",
		"Assessment.Scope.OutScope",
		"AssessmentStage",
	)
	if err != nil {
		return nil, fmt.Errorf("task with ID %d not found: %w", id, err)
	}
	return &task, nil
}

func (s *TaskService) GetAllTasks() ([]*entity.Task, error) {
	var tasks []*entity.Task

	err := s.repo.GetAll(&tasks,
		"Assessment",
		"Assessment.Scope",
		"Assessment.Scope.InScope",
		"Assessment.Scope.OutScope",
		"AssessmentStage",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %w", err)
	}
	return tasks, nil
}

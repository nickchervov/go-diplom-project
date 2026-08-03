package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
)

func (s *SchedulerService) GetTask(ctx context.Context, input dto.GetTaskInput) (dto.GetTaskOutput, error) {
	if input.Id == "" {
		return dto.GetTaskOutput{}, fmt.Errorf("id must be not empty: %w", domain.ErrIncorrectId)
	}
	id, err := strconv.Atoi(input.Id)
	if err != nil {
		return dto.GetTaskOutput{}, fmt.Errorf("id must be a number: %w", domain.ErrIncorrectId)
	}
	if id < 0 {
		return dto.GetTaskOutput{}, fmt.Errorf("id must be more then zero: %w", domain.ErrIncorrectId)
	}

	task, err := s.repo.GetTask(ctx, input.Id)
	if err != nil {
		return dto.GetTaskOutput{}, fmt.Errorf("getting task: %w", err)
	}

	output := dto.GetTaskOutput{
		Id:      task.Id,
		Date:    task.Date,
		Title:   task.Title,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}

	return output, nil
}

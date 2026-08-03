package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
)

func (s *SchedulerService) DeleteTask(ctx context.Context, input dto.DeleteTaskInput) error {
	if input.Id == "" {
		return fmt.Errorf("id must be not empty: %w", domain.ErrIncorrectId)
	}
	id, err := strconv.Atoi(input.Id)
	if err != nil {
		return fmt.Errorf("id must be a number: %w", domain.ErrIncorrectId)
	}
	if id < 0 {
		return fmt.Errorf("id must be more then zero: %w", domain.ErrIncorrectId)
	}

	if err := s.repo.DeleteTask(ctx, input.Id); err != nil {
		return fmt.Errorf("deleting task: %w", err)
	}

	return nil
}

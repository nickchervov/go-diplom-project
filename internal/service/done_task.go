package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/nextdate"
)

func (s *SchedulerService) DoneTask(ctx context.Context, input dto.DoneTaskInput) error {
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

	task, err := s.repo.GetTask(ctx, input.Id)
	if err != nil {
		return fmt.Errorf("getting task: %w", err)
	}

	if task.Repeat == "" {
		if err := s.repo.DeleteTask(ctx, input.Id); err != nil {
			return fmt.Errorf("deleting task: %w", err)
		}
		return nil
	}

	newDate, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return fmt.Errorf("getting next date: %w", err)
	}

	if err := s.repo.UpdateDate(ctx, input.Id, newDate); err != nil {
		return fmt.Errorf("updating date: %w", err)
	}

	return nil
}

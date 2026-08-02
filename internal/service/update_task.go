package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/nextdate"
)

func (s *SchedulerService) UpdateTask(ctx context.Context, input dto.UpdateTaskInput) error {
	if input.Title == "" {
		return domain.ErrEmptyTitle
	}
	if input.Date == "" {
		input.Date = time.Now().Format("20060102")
	}
	_, err := time.Parse("20060102", input.Date)
	if err != nil {
		return domain.ErrIncorrectDate
	}

	now := time.Now().Format("20060102")

	if input.Date < now && input.Repeat == "" {
		input.Date = now
	}
	if input.Date < now && input.Repeat != "" {
		nextDate, err := nextdate.NextDate(time.Now(), input.Date, input.Repeat)
		if err != nil {
			return fmt.Errorf("generating next date: %w", err)
		}
		input.Date = nextDate
	}

	task := domain.NewTask(input.Date, input.Title, input.Comment, input.Repeat)

	if err := s.repo.UpdateTask(ctx, input.Id, task); err != nil {
		return fmt.Errorf("updating task: %w", err)
	}
	return nil
}

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/nextdate"
)

func (s *SchedulerService) AddTask(ctx context.Context, input dto.CreateTaskInput) (dto.CreateTaskOutput, error) {
	if input.Title == "" {
		return dto.CreateTaskOutput{}, fmt.Errorf("%w", domain.ErrEmptyTitle)
	}
	if input.Date == "" {
		input.Date = time.Now().Format("20060102")
	}

	_, err := time.Parse("20060102", input.Date)
	if err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("%w", domain.ErrIncorrectDate)
	}

	now := time.Now().Format("20060102")

	if input.Date < now && input.Repeat == "" {
		input.Date = now
	}
	if input.Date < now && input.Repeat != "" {
		nextDate, err := nextdate.NextDate(time.Now(), input.Date, input.Repeat)
		if err != nil {
			return dto.CreateTaskOutput{}, fmt.Errorf("generating next date: %w", err)
		}
		input.Date = nextDate
	}

	task := domain.NewTask(input.Date, input.Title, input.Comment, input.Repeat)

	id, err := s.repo.AddTask(ctx, task)
	if err != nil {
		return dto.CreateTaskOutput{}, fmt.Errorf("adding task: %w", err)
	}

	output := dto.CreateTaskOutput{
		Id: fmt.Sprint(id),
	}
	return output, nil
}

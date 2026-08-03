package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/dto"
)

func isDate(search string) bool {
	_, err := time.Parse("02.01.2006", search)
	return err == nil
}

func (s *SchedulerService) GetTasks(ctx context.Context, search string, limit int) (dto.GetTasksOutput, error) {
	switch {
	case search != "" && !isDate(search):
		tasks, err := s.repo.GetTasksByTitleOrComment(ctx, search, limit)
		if err != nil {
			return dto.GetTasksOutput{}, fmt.Errorf("getting tasks by title or comment: %w", err)
		}
		return dto.GetTasksOutput{Tasks: tasks}, nil

	case search != "" && isDate(search):
		date, err := time.Parse("02.01.2006", search)
		if err != nil {
			return dto.GetTasksOutput{}, fmt.Errorf("parse searched date: %w", err)
		}
		tasks, err := s.repo.GetTasksByDate(ctx, date.Format("20060102"), limit)
		if err != nil {
			return dto.GetTasksOutput{}, fmt.Errorf("getting tasks by date: %w", err)
		}
		return dto.GetTasksOutput{Tasks: tasks}, nil

	default:
		tasks, err := s.repo.GetTasks(ctx, limit)
		if err != nil {
			return dto.GetTasksOutput{}, fmt.Errorf("getting tasks: %w", err)
		}
		return dto.GetTasksOutput{Tasks: tasks}, nil
	}
}

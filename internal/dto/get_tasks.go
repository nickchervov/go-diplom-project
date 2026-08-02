package dto

import (
	"github.com/nickchervov/go-diplom-project/internal/domain"
)

type GetTasksOutput struct {
	Tasks []domain.Task `json:"tasks"`
}

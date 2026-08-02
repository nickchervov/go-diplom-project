package domain

import "errors"

var (
	ErrEmptyTitle          = errors.New("title must be not empty")
	ErrIncorrectDate       = errors.New("incorrect date")
	ErrIncorrectRepeatRule = errors.New("incorrect repeat rule")
	ErrIncorrectId         = errors.New("incorrect id")
	ErrTaskNotFound        = errors.New("task not found")
	ErrIncorrectPassword   = errors.New("incorrect password")
)

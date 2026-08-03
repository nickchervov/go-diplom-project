package handler

import "github.com/nickchervov/go-diplom-project/internal/service"

type Handler struct {
	svc *service.SchedulerService
}

func NewHandler(svc *service.SchedulerService) *Handler {
	return &Handler{svc: svc}
}

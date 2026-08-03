package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	input := dto.GetTaskInput{
		Id: r.FormValue("id"),
	}

	task, err := h.svc.GetTask(r.Context(), input)
	if err != nil {
		var targetErr *domain.DomainError
		if errors.As(err, &targetErr) {
			render.Json(w, targetErr.Code, map[string]string{"error": targetErr.Error()})
			return
		}
		log.Printf("internal error get task: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusOK, task)
}

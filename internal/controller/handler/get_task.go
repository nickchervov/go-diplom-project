package handler

import (
	"errors"
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
		if errors.Is(err, domain.ErrIncorrectId) {
			render.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrTaskNotFound) {
			render.Json(w, http.StatusNotFound, map[string]string{"error": domain.ErrTaskNotFound.Error()})
			return
		}
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusOK, task)
}

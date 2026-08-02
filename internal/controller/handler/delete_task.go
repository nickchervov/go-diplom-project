package handler

import (
	"errors"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	input := dto.DeleteTaskInput{Id: r.FormValue("id")}

	if err := h.svc.DeleteTask(r.Context(), input); err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			render.Json(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrIncorrectId) {
			render.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	render.Json(w, http.StatusOK, dto.DeleteTaskOutput{})
}

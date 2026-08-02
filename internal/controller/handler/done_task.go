package handler

import (
	"errors"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) DoneTask(w http.ResponseWriter, r *http.Request) {
	input := dto.DoneTaskInput{Id: r.FormValue("id")}

	if err := h.svc.DoneTask(r.Context(), input); err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			render.Json(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		if errors.Is(err, domain.ErrIncorrectId) || errors.Is(err, domain.ErrIncorrectDate) || errors.Is(err, domain.ErrIncorrectRepeatRule) {
			render.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	render.Json(w, http.StatusOK, dto.DoneTaskOutput{})
}

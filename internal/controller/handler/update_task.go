package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "decoding request body"})
		return
	}

	if err := h.svc.UpdateTask(r.Context(), input); err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			render.Json(w, http.StatusNotFound, map[string]string{"error": domain.ErrTaskNotFound.Error()})
			return
		}
		if errors.Is(err, domain.ErrEmptyTitle) || errors.Is(err, domain.ErrIncorrectDate) || errors.Is(err, domain.ErrIncorrectRepeatRule) {
			render.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusOK, dto.UpdateTaskOutput{})
}

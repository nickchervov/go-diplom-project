package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.Json(w, http.StatusBadRequest, map[string]string{"error": "decoding request body"})
		return
	}

	if err := h.svc.UpdateTask(r.Context(), input); err != nil {
		var targetErr *domain.DomainError
		if errors.As(err, &targetErr) {
			render.Json(w, targetErr.Code, map[string]string{"error": targetErr.Error()})
			return
		}
		log.Printf("internal error update task: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusOK, dto.UpdateTaskOutput{})
}

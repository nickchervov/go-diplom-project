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

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.Json(w, http.StatusBadRequest, map[string]string{"error": "decoding request body"})
		return
	}

	output, err := h.svc.AddTask(r.Context(), input)
	if err != nil {
		var targetErr *domain.DomainError
		if errors.As(err, &targetErr) {
			render.Json(w, targetErr.Code, map[string]string{"error": targetErr.Error()})
			return
		}
		log.Printf("internal error add task: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusCreated, output)
}

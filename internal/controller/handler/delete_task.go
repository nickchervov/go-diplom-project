package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	input := dto.DeleteTaskInput{Id: r.FormValue("id")}

	if err := h.svc.DeleteTask(r.Context(), input); err != nil {
		var targetErr *domain.DomainError
		if errors.As(err, &targetErr) {
			render.Json(w, targetErr.Code, map[string]string{"error": targetErr.Error()})
			return
		}
		log.Printf("internal error delete task: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	render.Json(w, http.StatusOK, dto.DeleteTaskOutput{})
}

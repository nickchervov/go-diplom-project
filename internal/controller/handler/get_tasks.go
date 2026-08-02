package handler

import (
	"net/http"

	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := h.svc.GetTasks(r.Context(), search, 50)
	if err != nil {
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	render.Json(w, http.StatusOK, tasks)
}

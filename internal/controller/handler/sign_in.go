package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/internal/dto"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input dto.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "decoding request body"})
		return
	}

	output, err := h.svc.SignIn(r.Context(), input)
	if err != nil {
		if errors.Is(err, domain.ErrIncorrectPassword) {
			render.Json(w, http.StatusUnauthorized, map[string]string{"error": domain.ErrIncorrectPassword.Error()})
			return
		}
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}

	render.Json(w, http.StatusOK, output)
}

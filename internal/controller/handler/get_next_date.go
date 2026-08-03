package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/nickchervov/go-diplom-project/internal/domain"
	"github.com/nickchervov/go-diplom-project/pkg/nextdate"
	"github.com/nickchervov/go-diplom-project/pkg/render"
)

func (h *Handler) GetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if now == "" {
		now = time.Now().Format("20060102")
	}

	dateNow, err := time.Parse("20060102", now)
	if err != nil {
		log.Printf("internal error get next date: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "parsing now date"})
		return
	}
	nextDate, err := nextdate.NextDate(dateNow, date, repeat)
	if err != nil {
		var targetErr *domain.DomainError
		if errors.As(err, &targetErr) {
			render.Json(w, targetErr.Code, map[string]string{"error": targetErr.Error()})
			return
		}
		log.Printf("internal error get next date: %v\n", err)
		render.Json(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}

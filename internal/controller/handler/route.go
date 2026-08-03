package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/nickchervov/go-diplom-project/internal/service"
)

func SetRoutes(svc *service.SchedulerService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := NewHandler(svc)

	r.Handle("/*", http.FileServer(http.Dir("web")))

	r.Get("/api/nextdate", h.GetNextDate)

	r.With(auth).Get("/api/tasks", h.GetTasks)

	r.With(auth).Post("/api/task", h.AddTask)
	r.With(auth).Get("/api/task", h.GetTask)
	r.With(auth).Put("/api/task", h.UpdateTask)
	r.With(auth).Delete("/api/task", h.DeleteTask)

	r.With(auth).Post("/api/task/done", h.DoneTask)

	r.Post("/api/signin", h.SignIn)

	return r
}

package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jseiser/justin-go/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/health", loadhealthRoutes)
	router.Route("/todo", loadtodoRoutes)

	return router
}

func loadhealthRoutes(router chi.Router) {
	healthHandler := &handler.Health{}

	router.Get("/live", healthHandler.Live)
	router.Get("/ready", healthHandler.Ready)
}

func loadtodoRoutes(router chi.Router) {
	todoHandler := &handler.Todo{}

	router.Post("/", todoHandler.Create)
	router.Get("/", todoHandler.List)
	router.Get("/{id}", todoHandler.GetByID)
	router.Put("/{id}", todoHandler.UpdateByID)
	router.Delete("/{id}", todoHandler.DeleteByID)
}

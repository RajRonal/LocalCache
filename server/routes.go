package server

import (
	"InMemoryCache/providers/cacheProvider"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Cache *cacheProvider.Caches
	http.Server
}

func NewServer() *Server {
	return &Server{
		Cache: cacheProvider.InitCache(),
	}
}

func (s *Server) SetupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Post("/create", s.InsertData)
		api.Get("/", s.GetData)
	})

	return router

}
func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc.SetupRoutes())
}

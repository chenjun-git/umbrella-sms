package handler

import (
	"github.com/go-chi/chi"

	"github.com/chenjun-git/umbrella-sms/handler/v1.0"
)

func RegisterSmsRouter() *chi.Mux {
	router := chi.NewRouter()
	v1_0.RegisterRouter(router)
	return router
}

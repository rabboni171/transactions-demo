package router

import (
	"github.com/gorilla/mux"
	"github.com/rabboni171/transactions-demo/internal/handler"
)

// InitRoutes инициализирует все маршруты для приложения
func InitRoutes(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()

	// Регистрируем маршруты
	router.HandleFunc("/transfer", handler.Transfer).Methods("POST")

	return router
}

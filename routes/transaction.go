package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.POST("/notification", h.Notification)
	e.POST("/createtransaction", h.CreateTransaction)
	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/{id}", h.GetTransaction)

	// r.HandleFunc("/createtransaction", h.CreateTransaction).Methods("POST")
	// r.HandleFunc("/notification", h.Notification).Methods("POST")
	// r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	// r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
}

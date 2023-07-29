package main

import (
	"curso-go/internal/entity"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", Order)

	http.ListenAndServe(":8888", r)

}

func Order(w http.ResponseWriter, r *http.Request) {
	order, err := entity.NewOrder("123", 50, 10)
	order.CalculateFinalPrice()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

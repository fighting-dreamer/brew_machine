package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/server/handler"
)

func handleIngredientRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	ingredientHandler := handler.NewIngredientHandler(dependencies)
	router.HandleFunc("/v1/ingredient/add", ingredientHandler.AddNewIngredientAPI).Methods(http.MethodPost)
	router.HandleFunc("/v1/ingredient/refill", ingredientHandler.RefillIngredientAPI).Methods(http.MethodPut)
	router.HandleFunc("/v1/ingredient/available", ingredientHandler.IsIngredientAvailableAPI).Methods(http.MethodGet)
}

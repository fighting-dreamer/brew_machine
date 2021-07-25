package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/server/handler"
)

func handleBeverageRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	beverageHandler := handler.NewBeverageHandler(dependencies)
	router.HandleFunc("/v1/beverage/add", beverageHandler.AddNewBeverageAPI).Methods(http.MethodPost)
}

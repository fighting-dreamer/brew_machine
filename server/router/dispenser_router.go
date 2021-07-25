package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/server/handler"
)

func handleDispenserRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	dispenserHandler := handler.NewDispenserHandler(dependencies)
	router.HandleFunc("/v1/dispenser/make", dispenserHandler.MakeBeverageAPI).Methods(http.MethodGet)
}

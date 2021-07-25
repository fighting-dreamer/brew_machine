package router

import (
	"github.com/gorilla/mux"
	"nipun.io/brew_machine/appcontext"
)

func Router(dependencies *appcontext.Instance) *mux.Router {
	router := mux.NewRouter()

	handleSystemRoutes(dependencies, router)

	return router
}
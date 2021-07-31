package router

import (
	"net/http"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/server/handler"

	"github.com/gorilla/mux"
)

func handleSystemRoutes(dependencies *appcontext.Instance, router *mux.Router) {
	router.HandleFunc("/ping", handler.PingHandler).
		Methods(http.MethodGet)
	//LockHandler := handler.NewLockHandler(dependencies)
	//router.HandleFunc("/lock_state", LockHandler.GetCurrentLocksStateAPI).Methods(http.MethodGet)
}

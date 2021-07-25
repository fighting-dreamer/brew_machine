package server

import (
	"github.com/urfave/negroni"
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/config"
	"nipun.io/brew_machine/logger"
	r "nipun.io/brew_machine/server/router"
	"strconv"
)

func StartApiServer(dependencies *appcontext.Instance) {
	logger.Logger.Info().Msg("Starting API server")

	router := r.Router(dependencies)
	n := negroni.New(negroni.NewRecovery())
	n.UseHandlerFunc(router.ServeHTTP)

	portInfo := ":" + strconv.Itoa(config.AppPort())
	n.Run(portInfo)
}

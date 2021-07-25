package main

import (
	"nipun.io/brew_machine/appcontext"
	"nipun.io/brew_machine/server"
)

func main() {
	appcontext.Init()

	server.StartApiServer(appcontext.AppDependencies)
}

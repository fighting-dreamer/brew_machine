package appcontext

import (
	"nipun.io/brew_machine/config"
	"nipun.io/brew_machine/logger"
)

func Init() {
	config.Load()
	logger.SetupLogger()
	LoadDependencies()
}

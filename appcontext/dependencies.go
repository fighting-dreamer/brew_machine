package appcontext

import (
	"nipun.io/brew_machine/repository"
	local_repo "nipun.io/brew_machine/repository/local"
)

type Instance struct {
	BeverageRepository   repository.BeverageRepository
	IngredientRepository repository.IngredientRepository
}

var AppDependencies *Instance

func LoadDependencies() {
	AppDependencies = &Instance{}
}

func (dependencies *Instance) addBeverageRepository() {
	dependencies.BeverageRepository = local_repo.NewInMemoryBeverageRepository()
}

func (dependencies *Instance) addIngredientRepository() {
	dependencies.IngredientRepository = local_repo.NewInMemoryIngredientRepository()
}

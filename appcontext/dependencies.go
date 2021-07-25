package appcontext

import (
	"nipun.io/brew_machine/repository"
	local_repo "nipun.io/brew_machine/repository/local"
	"nipun.io/brew_machine/service"
	local_service "nipun.io/brew_machine/service/local"
)

type Instance struct {
	BeverageRepository   repository.BeverageRepository
	IngredientRepository repository.IngredientRepository
	BeverageManager      service.BeverageManager
	IngredientManager    service.IngredientManager
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

func (dependencies *Instance) addBeverageManager() {
	dependencies.BeverageManager = &local_service.LocalBeverageManager{
		BeverageRepository: dependencies.BeverageRepository,
	}
}

func (dependencies *Instance) addIngredientManager() {
	dependencies.IngredientManager = &local_service.LocalIngredientManager{
		IngredientRepository: dependencies.IngredientRepository,
	}
}

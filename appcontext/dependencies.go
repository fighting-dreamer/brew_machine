package appcontext

import (
	"nipun.io/brew_machine/config"
	"nipun.io/brew_machine/repository"
	local_repo "nipun.io/brew_machine/repository/local"
	"nipun.io/brew_machine/service"
	local_service "nipun.io/brew_machine/service/local"
	"sync"
)

type Instance struct {
	BeverageRepository     repository.BeverageRepository
	IngredientRepository   repository.IngredientRepository
	BeverageManager        service.BeverageManager
	IngredientManager      service.IngredientManager
	DispenserService       service.DispenserService
	TransactionLockManager service.TransactionLockManager
}

var AppDependencies *Instance

func LoadDependencies() {
	AppDependencies = &Instance{}
	AppDependencies.addTransactionLockManager()
	AppDependencies.addBeverageRepository()
	AppDependencies.addIngredientRepository()
	AppDependencies.addBeverageManager()
	AppDependencies.addIngredientManager()
	AppDependencies.addDispenserService()
}

func (dependencies *Instance) addTransactionLockManager() {
	dependencies.TransactionLockManager = &local_service.TransactionLockManager{
		KeeperState: sync.Map{},//map[string]string{},
		Keeper:      sync.Map{},//map[string]*local_service.CounterMutex{},
	}
}

func (dependencies *Instance) addBeverageRepository() {
	dependencies.BeverageRepository = local_repo.NewInMemoryBeverageRepository(dependencies.TransactionLockManager)
}

func (dependencies *Instance) addIngredientRepository() {
	dependencies.IngredientRepository = local_repo.NewInMemoryIngredientRepository(dependencies.TransactionLockManager)
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

func (dependencies *Instance) addDispenserService() {
	dependencies.DispenserService = &local_service.LocalDispenserService{
		BeverageRepository:     dependencies.BeverageRepository,
		IngredientRepository:   dependencies.IngredientRepository,
		OutletCnt:              config.OutletCnt(),
		TransactionLockManager: dependencies.TransactionLockManager,
	}
}

package repository

import (
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/service"
)

type InMemoryBeverageRepository struct {
	data                   map[string]*domain.Beverage
	transactionLockManager service.TransactionLockManager
}

var (
	InMemoryBeverageRepositoryLockKeys = []string{"InMemoryBeverageRepository", "data"}
)

func NewInMemoryBeverageRepository(TransactionLockManager service.TransactionLockManager) *InMemoryBeverageRepository {
	return &InMemoryBeverageRepository{
		data:                   map[string]*domain.Beverage{},
		transactionLockManager: TransactionLockManager,
	}
}

func (imbr *InMemoryBeverageRepository) AddNew(beverage domain.Beverage) error {
	defer imbr.transactionLockManager.ReleaseLock(InMemoryBeverageRepositoryLockKeys)
	imbr.transactionLockManager.AcquireLock(InMemoryBeverageRepositoryLockKeys)
	if imbr.data[beverage.Name] != nil {
		return nil
	}
	imbr.data[beverage.Name] = &beverage
	return nil
}

func (imbr *InMemoryBeverageRepository) Get(name string) (*domain.Beverage, error) {
	defer imbr.transactionLockManager.ReleaseLock(InMemoryBeverageRepositoryLockKeys)
	imbr.transactionLockManager.AcquireLock(InMemoryBeverageRepositoryLockKeys)
	return imbr.data[name], nil
}

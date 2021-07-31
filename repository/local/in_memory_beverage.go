package repository

import (
	"errors"
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

const (
	BeverageAlreadyExists = "BeverageAlreadyExists"
	BeverageDoesNotExists = "BeverageDoesNotExists"
)

func NewInMemoryBeverageRepository(TransactionLockManager service.TransactionLockManager) *InMemoryBeverageRepository {
	return &InMemoryBeverageRepository{
		data:                   map[string]*domain.Beverage{},
		transactionLockManager: TransactionLockManager,
	}
}

func (imbr *InMemoryBeverageRepository) AddNew(beverage domain.Beverage) error {
	imbr.transactionLockManager.AcquireLock(InMemoryBeverageRepositoryLockKeys)
	if imbr.data[beverage.Name] != nil {
		imbr.transactionLockManager.ReleaseLock(InMemoryBeverageRepositoryLockKeys)
		return errors.New(BeverageAlreadyExists)
	}
	imbr.data[beverage.Name] = &beverage
	imbr.transactionLockManager.ReleaseLock(InMemoryBeverageRepositoryLockKeys)
	return nil
}

func (imbr *InMemoryBeverageRepository) Get(name string) (*domain.Beverage, error) {
	imbr.transactionLockManager.AcquireLock(InMemoryBeverageRepositoryLockKeys)
	beverage := imbr.data[name]
	imbr.transactionLockManager.ReleaseLock(InMemoryBeverageRepositoryLockKeys)
	if beverage == nil {
		return nil, errors.New(BeverageDoesNotExists)
	}
	return beverage, nil
}

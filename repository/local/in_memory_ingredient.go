package repository

import (
	"errors"
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/service"
)

var (
	IngredientDoesNotExist     = errors.New("IngredientDoesNotExist")
	IngredientLessThanRequired = errors.New("IngredientLessThanRequired")

	InMemoryIngredientRepositoryLockKeys = []string{"InMemoryIngredientRepository", "data"}
)

type InMemoryIngredientRepository struct {
	data                   map[string]*domain.Ingredient
	transactionLockManager service.TransactionLockManager
}

func NewInMemoryIngredientRepository(TransactionLockManager service.TransactionLockManager) *InMemoryIngredientRepository {
	return &InMemoryIngredientRepository{
		data:                   map[string]*domain.Ingredient{},
		transactionLockManager: TransactionLockManager,
	}
}

func (imir *InMemoryIngredientRepository) AddNew(ingredient domain.Ingredient) error {
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[ingredient.Name] != nil {
		imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
		return nil
	}
	imir.data[ingredient.Name] = &ingredient
	imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	return nil
}

func (imir *InMemoryIngredientRepository) UpdateQuantity(name string, delta int) error {
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[name] == nil {
		imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
		return IngredientDoesNotExist
	}

	ingredient := imir.data[name]
	if (ingredient.AvailableQuantity + delta) < 0 {
		imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
		return IngredientLessThanRequired
	}

	ingredient.AvailableQuantity = ingredient.AvailableQuantity + delta
	imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	return nil
}

func (imir *InMemoryIngredientRepository) Get(name string) (*domain.Ingredient, error) {
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[name] == nil {
		imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
		return nil, IngredientDoesNotExist
	}

	imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	return imir.data[name], nil
}

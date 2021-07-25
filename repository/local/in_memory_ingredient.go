package repository

import (
	"errors"
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/service"
)

var (
	IngredientDoesNotExist = errors.New("IngredientDoesNotExist")

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
	defer imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[ingredient.Name] != nil {
		return nil
	}
	imir.data[ingredient.Name] = &ingredient
	return nil
}

func (imir *InMemoryIngredientRepository) UpdateQuantity(name string, delta int) error {
	defer imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[name] == nil {
		return IngredientDoesNotExist
	}
	ingredient := imir.data[name]
	ingredient.AvailableQuantity = ingredient.AvailableQuantity + delta
	return nil
}

func (imir *InMemoryIngredientRepository) Get(name string) (*domain.Ingredient, error) {
	defer imir.transactionLockManager.ReleaseLock(InMemoryIngredientRepositoryLockKeys)
	imir.transactionLockManager.AcquireLock(InMemoryIngredientRepositoryLockKeys)
	if imir.data[name] == nil {
		return nil, IngredientDoesNotExist
	}
	return imir.data[name], nil
}

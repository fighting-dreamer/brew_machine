package repository

import (
	"errors"
	"nipun.io/brew_machine/domain"
)

var (
	IngredientDoesNotExist = errors.New("IngredientDoesNotExist")
)

type InMemoryIngredientRepository struct {
	data map[string]*domain.Ingredient
}

func NewInMemoryIngredientRepository() *InMemoryIngredientRepository {
	return &InMemoryIngredientRepository{
		data: map[string]*domain.Ingredient{},
	}
}

func (imir *InMemoryIngredientRepository) AddNew(ingredient domain.Ingredient) error {
	if imir.data[ingredient.Name] != nil {
		return nil
	}
	imir.data[ingredient.Name] = &ingredient
	return nil
}

func (imir *InMemoryIngredientRepository) UpdateQuantity(name string, delta int) error {
	if imir.data[name] == nil {
		return IngredientDoesNotExist
	}
	ingredient := imir.data[name]
	ingredient.AvailableQuantity = ingredient.AvailableQuantity + delta
	return nil
}

func (imir *InMemoryIngredientRepository) Get(name string) (*domain.Ingredient, error) {
	if imir.data[name] == nil {
		return nil, IngredientDoesNotExist
	}
	return imir.data[name], nil
}

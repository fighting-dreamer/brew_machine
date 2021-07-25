package service

import (
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/repository"
)

type LocalIngredientManager struct {
	IngredientRepository repository.IngredientRepository
}

func (lim *LocalIngredientManager) AddNew(ingredient domain.Ingredient) error {
	return lim.IngredientRepository.AddNew(ingredient)
}

func (lim *LocalIngredientManager) UpdateQuantity(name string, delta int) error {
	return lim.IngredientRepository.UpdateQuantity(name, delta)
}

func (lim *LocalIngredientManager) Get(name string) (*domain.Ingredient, error) {
	return lim.IngredientRepository.Get(name)
}

func (lim *LocalIngredientManager) IsAvailable(name string) (bool, error) {
	ingredient, err := lim.Get(name)
	if err != nil {
		return false, err
	}
	if ingredient != nil && ingredient.AvailableQuantity > 0 {
		return true, nil
	}
	return false, nil
}

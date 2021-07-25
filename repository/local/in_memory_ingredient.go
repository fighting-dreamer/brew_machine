package repository

import "nipun.io/brew_machine/domain"

type InMemoryIngredientRepository struct {
	data map[string]*domain.Ingredient
}

func NewInMemoryIngredientRepository() *InMemoryIngredientRepository {
	return &InMemoryIngredientRepository{
		data: map[string]*domain.Ingredient{},
	}
}

func (imir *InMemoryIngredientRepository) AddNew(ingredient domain.Ingredient) error {
	imir.data[ingredient.Name] = &ingredient
	return nil
}
func (imir *InMemoryIngredientRepository) UpdateQuantity(name string, delta int) error {
	ingredient := imir.data[name]
	ingredient.AvailableQuantity = ingredient.AvailableQuantity + delta
	return nil
}
func (imir *InMemoryIngredientRepository) Get(name string) (*domain.Ingredient, error) {
	return imir.data[name], nil
}

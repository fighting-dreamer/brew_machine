package repository

import "nipun.io/brew_machine/domain"

type IngredientRepository interface {
	AddNew(ingredient domain.Ingredient) error
	UpdateQuantity(name string, delta int) error
	Get(name string) (*domain.Ingredient, error)
}

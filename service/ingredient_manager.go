package service

import "nipun.io/brew_machine/domain"

type IngredientManager interface {
	AddNew(ingredient domain.Ingredient) error
	UpdateQuantity(name string, delta int) error
	Get(name string) (*domain.Ingredient, error)
	IsAvailable(name string) (bool, error)
}

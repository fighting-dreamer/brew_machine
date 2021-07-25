package repository

import "nipun.io/brew_machine/domain"

type BeverageRepository interface {
	AddNew(beverage domain.Beverage) error
	Get(name string) (*domain.Beverage, error)
}

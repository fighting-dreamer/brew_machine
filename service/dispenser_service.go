package service

import "nipun.io/brew_machine/domain"

type DispenserService interface {
	MakeBeverage(name string, outlet int) (domain.Beverage, error)
}

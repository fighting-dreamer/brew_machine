package service

import "nipun.io/brew_machine/domain"

type BeverageManager interface {
	AddNew(beverage domain.Beverage) error
}

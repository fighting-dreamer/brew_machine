package service

import (
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/repository"
)

type LocalBeverageManager struct {
	BeverageRepository repository.BeverageRepository
}

func (lbm *LocalBeverageManager) AddNew(beverage domain.Beverage) error {
	return lbm.BeverageRepository.AddNew(beverage)
}

package repository

import "nipun.io/brew_machine/domain"

type InMemoryBeverageRepository struct {
	data map[string]*domain.Beverage
}

func NewInMemoryBeverageRepository() *InMemoryBeverageRepository {
	return &InMemoryBeverageRepository{
		data: map[string]*domain.Beverage{},
	}
}

func (imbr *InMemoryBeverageRepository) AddNew(beverage domain.Beverage) error {
	if imbr.data[beverage.Name] != nil {
		return nil
	}
	imbr.data[beverage.Name] = &beverage
	return nil
}

func (imbr *InMemoryBeverageRepository) Get(name string) (*domain.Beverage, error) {
	return imbr.data[name], nil
}
